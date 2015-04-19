package hf

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"

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
	WriteTo(b Writer, parent Node, opt RenderOptions)
}

type Writer interface {
	io.ByteWriter
	io.Writer
	WriteString(string) (int, error)
}

// Text
type HTMLBytes []byte

var _ Node = HTMLBytes{}

func (h HTMLBytes) String() string {
	return string(h)
}

func (h HTMLBytes) WriteTo(b Writer, parent Node, opt RenderOptions) {
	b.Write([]byte(h))
}

func NodeToHTMLBytes(el Node, opt RenderOptions) HTMLBytes {
	var b villa.ByteSlice
	el.WriteTo(&b, nil, opt)
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
	Type       TagType
	attributes Attributes
	classes    []HTMLBytes
}

var _ Node = (*Void)(nil)

func (v *Void) WriteTo(b Writer, parent Node, opt RenderOptions) {
	b.WriteByte('<')
	b.Write(TagBytes[v.Type])

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

func (v *Void) Attr(name string, value string) *Void {
	if len(name) == 0 {
		// ignore empty name
		return v
	}
	name = strings.ToLower(name)

	// TODO ignore invalid attribute name
	if name == "class" {
		classes := bytes.Split([]byte(attrEscaper(value)), []byte{' '})

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

	if v.attributes == nil {
		v.attributes = make(Attributes)
	}
	v.attributes[name] = attrEscaper(value)
	return v
}

func findStrInArr(s HTMLBytes, arr []HTMLBytes) int {
	for i, el := range arr {
		if bytes.Equal(el, s) {
			return i
		}
	}

	return -1
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

func (t *Element) Name() HTMLBytes {
	return HTMLBytes(TagBytes[t.Type])
}

func (t *Element) Children() []Element {
	return t.Children()
}

func (e *Element) Attr(name string, value string) *Element {
	e.Void.Attr(name, value)
	return e
}

func (t *Element) Child(el ...Node) *Element {
	t.children = append(t.children, el...)

	return t
}

func (t *Element) T(txt string) *Element {
	return t.Child(T(txt))
}

func newNewLine(e *Element) bool {
	switch e.Type {
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

func (e *Element) WriteTo(b Writer, parent Node, opt RenderOptions) {
	// TODO omit and indent
	// Write the open tag including attributes
	e.Void.WriteTo(b, parent, opt)

	if newNewLine(e) {
		b.WriteByte('\n')
	}

	for _, child := range e.children {
		child.WriteTo(b, e, opt)
	}
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

var _ Node = (*Html)(nil)

func (h Html) WriteTo(b Writer, parent Node, opt RenderOptions) {
	doctypeBytes.WriteTo(b, parent, opt)
	b.WriteByte('\n')

	h.Element.WriteTo(b, h, opt)
}

func (h Html) Head() *Element {
	return h.children[0].(*Element)
}

func (h Html) Body() *Element {
	return h.children[1].(*Element)
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
