package hf

import (
	"strconv"

	"github.com/daviddengcn/go-villa"
)

// HTML Entities
var (
	// HTML Character Entity copyright
	COPY = HTMLNode("&copy;")
	// HTML Character Entity ampersand
	AMP = HTMLNode("&amp;")
	// HTML Character Entity less than
	LT = HTMLNode("&lt;")
	// HTML Character Entity greater than
	GT = HTMLNode("&gt;")
	// HTML Character Entity cent
	CENT = HTMLNode("&cent;")
	// HTML Character Entity pound
	POUND = HTMLNode("&pound;")
	// HTML Character Entity yen
	YEN = HTMLNode("&yen;")
	// HTML Character Entity euro
	EURO = HTMLNode("&euro;")
	// HTML Character Entity registered trademark
	REG = HTMLNode("&reg;")
	// HTML Character Entity non-breaking space
	NBSP  = HTMLNode("&nbsp;")
	TIMES = HTMLNode("&times;")
	LAQUO = HTMLNode("&laquo;")
	RAQUO = HTMLNode("&raquo;")

	//TODO define all entities
)

var nePrefix = HTMLNode("&#")

// NumEnt returns HTMLBytes a numerical entity.
func NumEnt(num int) HTMLNode {
	var b villa.ByteSlice

	b.Write([]byte(nePrefix))
	b.WriteString(strconv.Itoa(num))
	b.WriteByte(';')

	return HTMLNode(b)
}
