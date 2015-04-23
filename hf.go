package hf

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
	"strconv"

	. "github.com/daviddengcn/go-html-frame/htmldef"
	"github.com/daviddengcn/go-villa"
)

type RenderOptions struct {
	Ident       HTMLBytes
	DisableOmit bool
	// Sort attributes names before export.
	// This is useful for testing because otherwise the exported attributes are random.
	SortAttr bool
}

var Default = RenderOptions{}

type Node interface {
	Type() TagType
	WriteTo(b Writer, opt RenderOptions, parent *Element, childIndex int)
}

type Writer interface {
	io.ByteWriter
	io.Writer
	WriteString(string) (int, error)
}

type URL []byte

// Text
type HTMLBytes []byte

var _ Node = HTMLBytes{}

func (h HTMLBytes) String() string {
	return string(h)
}

func (h HTMLBytes) WriteTo(b Writer, opt RenderOptions, parent *Element, childIndex int) {
	b.Write([]byte(h))
}

func (h HTMLBytes) Type() TagType {
	return TextType
}

func NodeToHTMLBytes(el Node, opt RenderOptions) HTMLBytes {
	var b villa.ByteSlice
	el.WriteTo(&b, opt, nil, 0)
	return HTMLBytes(b)
}

type Attributes map[string]HTMLBytes

func (attrs Attributes) WriteTo(b Writer, sortAttr bool) {
	if sortAttr {
		names := make([]string, 0, len(attrs))
		for name := range attrs {
			names = append(names, name)
		}
		sort.Strings(names)

		for _, name := range names {
			value := attrs[name]

			b.WriteByte(' ')
			b.WriteString(name)
			if len(value) > 0 {
				b.WriteByte('=')
				b.WriteByte('"')
				b.Write([]byte(value))
			}
			b.WriteByte('"')
		}
		return
	}

	for name, value := range attrs {
		b.WriteByte(' ')
		b.WriteString(name)
		if len(value) > 0 {
			b.WriteByte('=')
			b.WriteByte('"')
			b.Write([]byte(value))
		}
		b.WriteByte('"')
	}
}

// An HTML void element
type Void struct {
	tagType    TagType
	attributes Attributes
	classes    []HTMLBytes
}

var _ Node = (*Void)(nil)

func (v *Void) Type() TagType {
	return v.tagType
}

// Implementation of Node.WriteTo. This will be called to generate open tags of both void and normal elements
func (v *Void) WriteTo(b Writer, opt RenderOptions, parent *Element, childIndex int) {
	b.WriteByte('<')
	b.Write(TagBytes[v.tagType])

	if len(v.classes) > 0 {
		b.WriteString(` class="`)
		b.Write([]byte(v.classes[0]))
		for i, n := 1, len(v.classes); i < n; i++ {
			b.WriteByte(' ')
			b.Write([]byte(v.classes[i]))
		}
		b.WriteByte('"')
	}
	v.attributes.WriteTo(b, opt.SortAttr)

	b.WriteByte('>')
}

func (v *Void) Name() HTMLBytes {
	return HTMLBytes(TagBytes[v.tagType])
}

func (v *Void) Attr(name string, value string) *Void {
	return v.attrOfBytes(name, attrEscaper(value))
}

func (v *Void) AttrIfNotEmpty(name, value string) *Void {
	if value == "" {
		return v
	}
	return v.Attr(name, value)
}

func (v *Void) attrOfBytes(name string, value HTMLBytes) *Void {
	if len(name) == 0 {
		// ignore empty name
		return v
	}
	name = strings.ToLower(name)

	if name == "class" {
		classes := bytes.Split([]byte(value), []byte{' '})

		v.classes = make([]HTMLBytes, 0, len(classes))
		for _, class := range classes {
			if len(class) == 0 {
				// continuous spaces
				continue
			}
			v.classes = append(v.classes, HTMLBytes(class))
		}
		return v
	}

	// TODO ignore invalid attribute name
	if v.attributes == nil {
		v.attributes = make(Attributes)
	}
	v.attributes[name] = value
	return v
}

func (v *Void) Title(title string) {
	v.Attr("title", title)
}

func (v *Void) TabIndex(tablInex int) {
	v.Attr("tableindex", strconv.Itoa(tablInex))
}

func (v *Void) NonEmptyAttr(name, value string) *Void {
	if value == "" {
		return v
	}
	return v.Attr(name, value)
}

func (v *Void) AddClass(classes ...string) *Void {
	for _, cls := range classes {
		clsBytes := attrEscaper(cls)
		if findStrInArr(clsBytes, v.classes) >= 0 {
			continue
		}
		v.classes = append(v.classes, clsBytes)
	}

	return v
}

func (t *Void) DelClass(classes ...string) *Void {
	for _, cls := range classes {
		clsBytes := attrEscaper(cls)
		i := findStrInArr(clsBytes, t.classes)
		if i >= 0 {
			copy(t.classes[i:], t.classes[i+1:])
			t.classes = t.classes[:len(t.classes)-1]
		}
	}

	return t
}

func (t *Void) ID(id string) *Void {
	return t.Attr("id", id)
}

type Element struct {
	Void
	children []Node
}

var _ Node = (*Element)(nil)

func (t *Element) Children() []Element {
	return t.Children()
}

func (e *Element) Attr(name string, value string) *Element {
	e.Void.Attr(name, value)
	return e
}

func (e *Element) NonEmptyAttr(name string, value string) *Element {
	e.Void.NonEmptyAttr(name, value)
	return e
}

func (e *Element) Child(el ...Node) *Element {
	e.children = append(e.children, el...)

	return e
}

func (e *Element) ChildEls(els ...*Element) *Element {
	for _, el := range els {
		e.children = append(e.children, el)
	}
	return e
}

func (e *Element) ChildVoids(vs ...*Void) *Element {
	for _, v := range vs {
		e.children = append(e.children, v)
	}
	return e
}

func (e *Element) T(txt string) *Element {
	return e.Child(T(txt))
}

func shouldNewLine(e *Element) bool {
	switch e.tagType {
	case PRETag, TEXTAREATag:
		if len(e.children) == 0 {
			return false
		}

		t, ok := e.children[0].(HTMLBytes)
		if !ok {
			return false
		}
		if len(t) == 0 {
			return false
		}
		return t[0] == '\n'
	}
	return false
}

func (e *Element) canOmitStartTag(parent *Element, childIndex int) bool {
	if len(e.attributes) > 0 || len(e.classes) > 0 {
		return false
	}

	switch e.tagType {
	case HTMLTag, HEADTag:
		return true

	case BODYTag:
		if len(e.children) == 0 {
			return true
		}
		switch e.children[0].Type() {
		case TextType:
			return !startWithSpace(e.children[0].(HTMLBytes))

		case METATag, LINKTag, SCRIPTTag, TEMPLATETag:
			return false
		}
		return true

	case COLGROUPTag:
		if len(e.children) == 0 {
			return false
		}

		if e.children[0].Type() != COLTag {
			return false
		}

		if childIndex > 0 && parent != nil && parent.children[childIndex-1].Type() == COLGROUPTag {
			return false
		}

		return true

	case TBODYTag:
		if len(e.children) == 0 {
			return false
		}

		if e.children[0].Type() != TRTag {
			return false
		}

		if childIndex > 0 && parent != nil {
			switch parent.children[childIndex-1].Type() {
			case TBODYTag, THEADTag, TFOOTTag:
				return false
			}
		}

		return true
	}

	return false
}

var pOmittedAfter = []bool{
	ADDRESSTag:    true,
	ARTICLETag:    true,
	ASIDETag:      true,
	BLOCKQUOTETag: true,
	DIVTag:        true,
	DLTag:         true,
	FIELDSETTag:   true,
	FOOTERTag:     true,
	FORMTag:       true,
	H1Tag:         true,
	H2Tag:         true,
	H3Tag:         true,
	H4Tag:         true,
	H5Tag:         true,
	H6Tag:         true,
	HEADERTag:     true,
	HGROUPTag:     true,
	HRTag:         true,
	MAINTag:       true,
	NAVTag:        true,
	OLTag:         true,
	PTag:          true,
	PRETag:        true,
	SECTIONTag:    true,
	TABLETag:      true,
	ULTag:         true,
}

func (e *Element) canOmitEndTag(parent *Element, childIndex int) bool {
	switch tp := e.Type(); tp {
	case HTMLTag, HEADTag, BODYTag:
		return true

	case LITag:
		if parent == nil {
			return false
		}
		if childIndex == len(parent.children)-1 {
			return true
		}
		if parent.children[childIndex+1].Type() == LITag {
			return true
		}

	case DTTag, DDTag:
		if childIndex == len(parent.children)-1 {
			return tp == DDTag
		}
		switch parent.children[childIndex+1].Type() {
		case DTTag, DDTag:
			return true
		}

	case PTag:
		if childIndex == len(parent.children)-1 {
			return parent.Type() != ATag
		}

		nextTp := parent.children[childIndex+1].Type()
		if nextTp < 0 || int(nextTp) >= len(pOmittedAfter) {
			return false
		}
		return pOmittedAfter[nextTp]

	case RBTag, RTTag, RTCTag, RPTag:
		if childIndex == len(parent.children)-1 {
			return true
		}

		switch parent.children[childIndex+1].Type() {
		case RBTag, RTCTag, RPTag:
			return true
		case RTTag:
			return tp != RTCTag
		}

	case OPTGROUPTag, OPTIONTag:
		if childIndex == len(parent.children)-1 {
			return true
		}

		switch parent.children[childIndex+1].Type() {
		case OPTGROUPTag:
			return true
		case OPTIONTag:
			return tp == OPTIONTag
		}

	case COLGROUPTag:
		if childIndex == len(parent.children)-1 {
			return true
		}

		if parent.children[childIndex+1].Type() == TextType {
			return !startWithSpace(e.children[childIndex+1].(HTMLBytes))
		}

		return true

	case THEADTag:
		if childIndex == len(parent.children)-1 {
			return false
		}

		switch parent.children[childIndex+1].Type() {
		case TBODYTag, TFOOTTag:
			return true
		}

	case TBODYTag, TFOOTTag:
		if childIndex == len(parent.children)-1 {
			return true
		}

		switch parent.children[childIndex+1].Type() {
		case TBODYTag:
			return true

		case TFOOTTag:
			return tp == TBODYTag
		}

	case TRTag:
		if childIndex == len(parent.children)-1 {
			return true
		}

		return parent.children[childIndex+1].Type() == TRTag

	case THTag, TDTag:
		if childIndex == len(parent.children)-1 {
			return true
		}

		switch parent.children[childIndex+1].Type() {
		case THTag, TDTag:
			return true
		}
	}

	return false
}

func (e *Element) WriteTo(b Writer, opt RenderOptions, parent *Element, childIndex int) {
	// TODO omit and indent
	if opt.DisableOmit || !e.canOmitStartTag(parent, childIndex) {
		// Write the open tag including attributes
		e.Void.WriteTo(b, opt, parent, childIndex)
	}

	if shouldNewLine(e) {
		b.WriteByte('\n')
	}

	for i, child := range e.children {
		child.WriteTo(b, opt, e, i)
	}

	if !opt.DisableOmit && e.canOmitEndTag(parent, childIndex) {
		return
	}

	b.WriteByte('<')
	b.WriteByte('/')
	b.Write(TagBytes[e.tagType])
	b.WriteByte('>')
}

func T(text string) HTMLBytes {
	return htmlEscaper(text)
}

func Tf(format string, args ...interface{}) HTMLBytes {
	return htmlEscaper(fmt.Sprintf(format, args...))
}

var (
	doctypeBytes = HTMLBytes("<!DOCTYPE html>")
)

type Html struct {
	Element
}

var _ Node = Html{}

// Implementation of Node interface
func (h Html) WriteTo(b Writer, opt RenderOptions, parent *Element, childIndex int) {
	b.Write(doctypeBytes)
	b.WriteByte('\n')

	h.Element.WriteTo(b, opt, parent, childIndex)
}

// Implementation of Node interface
func (h Html) Type() TagType {
	return HTMLTag
}

func (h Html) Head() *Element {
	return h.children[0].(*Element)
}

func (h Html) Body() *Element {
	return h.children[1].(*Element)
}

func (h Html) Lang(lang string) Html {
	h.NonEmptyAttr("lang", lang)
	return h
}

func (h Html) Title(title string) Html {
	h.Head().Child(TITLE(title))
	return h
}

func (h Html) Favicon(href, tp string) Html {
	h.Head().Child(LINK(href, "shortcut icon").Attr("type", tp))
	return h
}

func (h Html) Css(href string) Html {
	h.Head().Child(LINK(href, "stylesheet").Attr("href", href).Attr("type", "text/css"))
	return h
}

func (h Html) SCRIPT(src string, content string) Html {
	h.Body().Child(SCRIPT(src, content))
	return h
}
