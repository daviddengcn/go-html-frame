package hf

import (
	"fmt"
)

func ExampleHtml() {
	h := HTML()
	fmt.Println(NodeToHTMLBytes(h, RenderOptions{SortAttr: true}))
	fmt.Println(NodeToHTMLBytes(h, RenderOptions{SortAttr: true}))
	// OUTPUT:
	// <!DOCTYPE html>
	// <html lang="en"><head><meta charset="utf-8"><meta content="initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, width=device-width" name="viewport"><body>
}
