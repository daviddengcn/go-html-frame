package hf

import (
	"strconv"

	"github.com/daviddengcn/go-villa"
)

// HTML Entities
var (
	// HTML Character Entity copyright
	COPY = HTMLBytes("&copy;")
	// HTML Character Entity ampersand
	AMP = HTMLBytes("&amp;")
	// HTML Character Entity less than
	LT = HTMLBytes("&lt;")
	// HTML Character Entity greater than
	GT = HTMLBytes("&gt;")
	// HTML Character Entity cent
	CENT = HTMLBytes("&cent;")
	// HTML Character Entity pound
	POUND = HTMLBytes("&pound;")
	// HTML Character Entity yen
	YEN = HTMLBytes("&yen;")
	// HTML Character Entity euro
	EURO = HTMLBytes("&euro;")
	// HTML Character Entity registered trademark
	REG = HTMLBytes("&reg;")
	// HTML Character Entity non-breaking space
	NBSP  = HTMLBytes("&nbsp;")
	TIMES = HTMLBytes("&times;")
	LAQUO = HTMLBytes("&laquo;")
	RAQUO = HTMLBytes("&raquo;")

	//TODO define all entities
)

var nePrefix = HTMLBytes("&#")

// NumEnt returns HTMLBytes a numerical entity.
func NumEnt(num int) HTMLBytes {
	var b villa.ByteSlice

	b.Write([]byte(nePrefix))
	b.WriteString(strconv.Itoa(num))
	b.WriteByte(';')

	return HTMLBytes(b)
}
