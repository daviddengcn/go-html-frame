package hf

import (
	. "github.com/daviddengcn/go-html-frame/htmldef"
)

// the singleton for all BR tags
var brTag = Void{
	tagType: BRTag,
}

func BR() *Void {
	return &brTag
}

func IMG(src string) *Void {
	return (&Void{
		tagType: IMGTag,
	}).Attr("src", src)
}

func LINK(href, rel string) *Void {
	return (&Void{
		tagType: LINKTag,
	}).Attr("href", href).Attr("rel", rel)
}

func META() *Void {
	return &Void{
		tagType: METATag,
	}
}

func PARAM(name, value string) *Void {
	return (&Void{
		tagType: PARAMTag,
	}).Attr("name", name).Attr("value", value)
}

func A(href string, children ...Node) *Element {
	t := (&Element{
		Void: Void{tagType: ATag},
	}).Child(children...)

	if href != "" {
		t.Attr("href", href)
	}

	return t
}

func HTML(lang string) Html {
	return Html{
		Element: Element{
			Void: Void {tagType: HTMLTag},
			children: []Node{
				HEAD(),
				BODY(),
			},
		},
	}.Lang(lang)
}

func HEAD() *Element {
	return &Element{
		Void: Void{tagType: HEADTag},
		children: []Node{
			META().Attr("charset", "utf-8"),
			META().Attr("name", "viewport").Attr("content", "initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, width=device-width"),
		},
	}
}

func BODY() *Element {
	return &Element{
		Void: Void{tagType: BODYTag},
	}
}

func TITLE(title string) *Element {
	return &Element{
		Void: Void{tagType: TITLETag},
		children: []Node{
			T(title),
		},
	}
}

func H1(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: H1Tag},
	}).Child(children...)
}
func H2(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: H2Tag},
	}).Child(children...)
}
func H3(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: H3Tag},
	}).Child(children...)
}
func H4(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: H4Tag},
	}).Child(children...)
}
func H5(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: H5Tag},
	}).Child(children...)
}
func H6(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: H6Tag},
	}).Child(children...)
}

func SMALL(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: SMALLTag},
	}).Child(children...)
}

func UL(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: ULTag},
	}).Child(children...)
}

func LI(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: LITag},
	}).Child(children...)
}

func NAV(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: NAVTag},
	}).Child(children...)
}

func FORM(method, action string, children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: FORMTag},
	}).Attr("method", method).Attr("action", action).Child(children...)
}

func INPUT(tp, name, value string) *Void {
	v := &Void{
		tagType: INPUTTag,
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
		Void: Void{tagType: BUTTONTag},
	}).Child(children...)
}

func SCRIPT(src string, content string) *Element {
	t := &Element{Void: Void{tagType: SCRIPTTag}}
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
		Void: Void{tagType: NOSCRIPTTag},
	}).Child(children...)
}

func SPAN(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: SPANTag},
	}).Child(children...)
}

func FOOTER(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: FOOTERTag},
	}).Child(children...)
}

func LABEL(For string, children ...Node) *Element {
	t := (&Element{
		Void: Void{tagType: LABELTag},
	}).Child(children...)
	if For != "" {
		t.Attr("for", For)
	}
	return t
}

func TEXTAREA(name, value string, children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: TEXTAREATag},
	}).Attr("name", name).Attr("value", value).Child(children...)
}

func P(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: PTag},
	}).Child(children...)
}

func B(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: BTag},
	}).Child(children...)
}

func OL(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: OLTag},
	}).Child(children...)
}

func PRE(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: PRETag},
	}).Child(children...)
}

func OBJECT(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: OBJECTTag},
	}).Child(children...)
}

func DD(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: DDTag},
	}).Child(children...)
}

func DIV(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: DIVTag},
	}).Child(children...)
}

func DL(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: DLTag},
	}).Child(children...)
}

func DT(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: DTTag},
	}).Child(children...)
}

func EMBED(attrs ...string) *Element {
	var t = Element{
		Void: Void{tagType: EMBEDTag},
	}

	for i := 1; i < len(attrs); i += 2 {
		t.Attr(attrs[i-1], attrs[i])
	}

	return &t
}
	
func RB(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: RBTag},
	}).Child(children...)
}

func RP(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: RPTag},
	}).Child(children...)
}

func RT(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: RTTag},
	}).Child(children...)
}

func RTC(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: RTCTag},
	}).Child(children...)
}

func RUBY(children ...Node) *Element {
	return (&Element{
		Void: Void{tagType: RUBYTag},
	}).Child(children...)
}
