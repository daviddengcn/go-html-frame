package hf

import (
	"fmt"
)

func ExampleHtml() {
	h := HTML("")
	fmt.Println(NodeToHTMLBytes(h, RenderOptions{DisableOmit: true, SortAttr: true}))
	// OUTPUT:
	// <!DOCTYPE html>
	// <html><head><meta charset="utf-8"><meta content="initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, width=device-width" name="viewport"></head><body></body></html>
}

func ExampleOmittedTags_li() {
	h := HTML("")
	body := h.Body()
	body.Child(UL(
		LI(T("Hello")),
		LI(T("World")),
	))
	fmt.Println(NodeToHTMLBytes(h, RenderOptions{SortAttr: true}))
	// OUTPUT:
	// <!DOCTYPE html>
	// <meta charset="utf-8"><meta content="initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, width=device-width" name="viewport"><ul><li>Hello<li>World</ul>
}

func ExampleOmittedTags_dtdd() {
	h := HTML("")
	body := h.Body()
	body.Child(DL(
		DT(T("Hello")),
		DT(T("Hello")),
		DD(T("World")),
		DD(T("World")),
		DT(T("Hello")),
		DD(T("World")),
	))
	fmt.Println(NodeToHTMLBytes(h, RenderOptions{SortAttr: true}))
	// OUTPUT:
	// <!DOCTYPE html>
	// <meta charset="utf-8"><meta content="initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, width=device-width" name="viewport"><dl><dt>Hello<dt>Hello<dd>World<dd>World<dt>Hello<dd>World</dl>
}

func ExampleOmittedTags_p() {
	h := HTML("")
	body := h.Body()
	body.Child(
		P(T("Hello")),
		DIV(T("Hello")),
		P(T("World")),
	)
	fmt.Println(NodeToHTMLBytes(h, RenderOptions{SortAttr: true}))
	
	// OUTPUT:
	// <!DOCTYPE html>
	// <meta charset="utf-8"><meta content="initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, width=device-width" name="viewport"><p>Hello<div>Hello</div><p>World
}

func ExampleOmittedTags_ruby() {
	h := HTML("")
	body := h.Body()
	body.Child(RUBY(
		T("中文"),
		RB(T("Hello")),
		RT(T("zhongwen")),
		RTC(T("World")),
		RP(T("World")),
	))
	fmt.Println(NodeToHTMLBytes(h, RenderOptions{SortAttr: true}))
	
	// OUTPUT:
	// <!DOCTYPE html>
	// <meta charset="utf-8"><meta content="initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, width=device-width" name="viewport"><ruby>中文<rb>Hello<rt>zhongwen<rtc>World<rp>World</ruby>
}