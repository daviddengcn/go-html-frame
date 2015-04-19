package hf

import (
	. "github.com/daviddengcn/go-html-frame/htmldef"
)

// the singleton for all BR tags
var brTag = Void{
	Type: BRTag,
}

func BR() *Void {
	return &brTag
}

func IMG(src string) *Void {
	return (&Void{
		Type: IMGTag,
	}).Attr("src", src)
}

func LINK(href, rel string) *Void {
	return (&Void{
		Type: LINKTag,
	}).Attr("href", href).Attr("rel", rel)
}

func META() *Void {
	return &Void{
		Type: METATag,
	}
}

func PARAM(name, value string) *Void {
	return (&Void{
		Type: PARAMTag,
	}).Attr("name", name).Attr("value", value)
}

func HTML() Html {
	return Html{
		Element: Element{
			Void: Void{
				Type: HTMLTag,
				attributes: map[string]HTMLBytes{
					"lang": HTMLBytes("en"),
				},
			},
			children: []Node{
				HEAD(),
				BODY(),
			},
		},
	}
}

func HEAD() *Element {
	return &Element{
		Void: Void{Type: HEADTag},
		children: []Node{
			META().Attr("charset", "utf-8"),
			META().Attr("name", "viewport").Attr("content", "initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, width=device-width"),
		},
	}
}

func BODY() *Element {
	return &Element{
		Void: Void{Type: BODYTag},
	}
}

func TITLE(title string) *Element {
	return &Element{
		Void: Void{Type: TITLETag},
		children: []Node{
			T(title),
		},
	}
}

func H1(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: H1Tag},
	}).Child(children...)
}
func H2(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: H2Tag},
	}).Child(children...)
}
func H3(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: H3Tag},
	}).Child(children...)
}
func H4(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: H4Tag},
	}).Child(children...)
}
func H5(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: H5Tag},
	}).Child(children...)
}
func H6(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: H6Tag},
	}).Child(children...)
}

func DIV(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: DIVTag},
	}).Child(children...)
}

func SMALL(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: SMALLTag},
	}).Child(children...)
}

func UL(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: ULTag},
	}).Child(children...)
}

func LI(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: LITag},
	}).Child(children...)
}

func NAV(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: NAVTag},
	}).Child(children...)
}

func A(href string, children ...Node) *Element {
	t := (&Element{
		Void: Void{Type: ATag},
	}).Child(children...)

	if href != "" {
		t.Attr("href", href)
	}

	return t
}

func FORM(method, action string, children ...Node) *Element {
	return (&Element{
		Void: Void{Type: FORMTag},
	}).Attr("method", method).Attr("action", action).Child(children...)
}

func INPUT(tp, name, value string) *Void {
	v := &Void{
		Type: INPUTTag,
	}

	if tp != "" {
		v.Attr("type", tp)
	}

	if name != "" {
		v.Attr("name", name)
	}

	if value != "" {
		v.Attr("value", value)
	}

	return v
}

func BUTTON(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: BUTTONTag},
	}).Child(children...)
}

func SCRIPT(src string, content string) *Element {
	t := &Element{Void: Void{Type: SCRIPTTag}}
	if src != "" {
		t.Attr("src", src)
	}

	if content != "" {
		t.Child(HTMLBytes(content))
	}

	return t
}

func NOSCRIPT(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: NOSCRIPTTag},
	}).Child(children...)
}

func SPAN(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: SPANTag},
	}).Child(children...)
}

func FOOTER(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: FOOTERTag},
	}).Child(children...)
}

func LABEL(For string, children ...Node) *Element {
	t := (&Element{
		Void: Void{Type: LABELTag},
	}).Child(children...)
	if For != "" {
		t.Attr("for", For)
	}
	return t
}

func TEXTAREA(name, value string, children ...Node) *Element {
	return (&Element{
		Void: Void{Type: TEXTAREATag},
	}).Attr("name", name).Attr("value", value).Child(children...)
}

func P(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: PTag},
	}).Child(children...)
}

func B(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: BTag},
	}).Child(children...)
}

func OL(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: OLTag},
	}).Child(children...)
}

func PRE(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: PRETag},
	}).Child(children...)
}

func OBJECT(children ...Node) *Element {
	return (&Element{
		Void: Void{Type: OBJECTTag},
	}).Child(children...)
}

func EMBED(attrs ...string) *Element {
	var t = Element{
		Void: Void{Type: EMBEDTag},
	}

	for i := 1; i < len(attrs); i += 2 {
		t.Attr(attrs[i-1], attrs[i])
	}

	return &t
}
