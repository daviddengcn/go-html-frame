package hf

import (
	"fmt"
	"html/template"
	"strings"
	"unicode/utf8"

	"github.com/golangplus/bytes"
)

func Url(format string, args ...interface{}) string {
	// FIXME
	return fmt.Sprintf(format, args...)
}

func Script(format string, args ...interface{}) template.HTML {
	// FIXME
	return template.HTML(fmt.Sprintf(format, args...))
}

// htmlReplacementTable contains the runes that need to be escaped
// inside a quoted attribute value or in a text node.
var htmlReplacementTable = []HTMLNode{
	// http://www.w3.org/TR/html5/syntax.html#attribute-value-(unquoted)-state
	// U+0000 NULL Parse error. Append a U+FFFD REPLACEMENT
	// CHARACTER character to the current attribute's value.
	// "
	// and similarly
	// http://www.w3.org/TR/html5/syntax.html#before-attribute-value-state
	0:    HTMLNode("\uFFFD"),
	0xA0: NBSP,
	'"':  HTMLNode("&quot;"),
	'&':  AMP,
	'<':  LT,
	'>':  GT,
}

// htmlReplacementTable contains the runes that need to be escaped
// inside a quoted attribute value or in a text node.
var attrReplacementTable = []HTMLNode{
	// http://www.w3.org/TR/html5/syntax.html#attribute-value-(unquoted)-state
	// U+0000 NULL Parse error. Append a U+FFFD REPLACEMENT
	// CHARACTER character to the current attribute's value.
	// "
	// and similarly
	// http://www.w3.org/TR/html5/syntax.html#before-attribute-value-state
	0:    HTMLNode("\uFFFD"),
	0xA0: NBSP,
	'"':  HTMLNode("&quot;"),
	'&':  AMP,
}

// attrEscaper escapes for inclusion in quoted attribute values.
func attrEscaper(attr string) HTMLNode {
	return htmlReplacer(attr, attrReplacementTable, true)
}

func attrNameEscape(name string) HTMLNode {
	return htmlReplacer(strings.ToLower(name), attrReplacementTable, true)
}

// htmlEscaper escapes for inclusion in HTML text.
func htmlEscaper(text string) HTMLNode {
	return htmlReplacer(text, htmlReplacementTable, true)
}

// RFC 3986: reserved
var isUrlUnreserved filterTableArr

func init() {
	isUrlUnreserved.SetRange('a', 'z')
	isUrlUnreserved.SetRange('A', 'Z')
	isUrlUnreserved.SetRange('0', '9')
	isUrlUnreserved.SetByString("*-._~")
}

// RFC 3986: gen-delims
var isUrlGenDelimis = filterTableArrFromString(":/?#[]@")

// RFC 3986: sub-delims
var isUrlSubDelims = filterTableArrFromString("!$&'()*+,;=")

var isUrlIpLiteralChars filterTableArr

func init() {
	isUrlIpLiteralChars = isUrlUnreserved
	isUrlIpLiteralChars.UnionEqual(isUrlSubDelims)
	isUrlIpLiteralChars[':'] = true
}

var isUrlRegNameChars filterTableArr

func init() {
	isUrlRegNameChars = isUrlUnreserved
	isUrlRegNameChars.UnionEqual(isUrlSubDelims)
}

var dec2hex = []byte("0123456789ABCDEF")

func queryEscape(s string) []byte {
	var b bytesp.Slice
	scanned := 0
	bs := []byte(s)
	inplaceChange := false
	for i, r := range bs {
		if !isUrlUnreserved[r] {
			if r == ' ' {
				bs[i] = '+'
				inplaceChange = true
			} else {
				b.Write(bs[scanned:i])
				b.WriteByte('%')
				b.WriteByte(dec2hex[r/0x10])
				b.WriteByte(dec2hex[r%0x10])
				scanned = i + 1
			}
		}
	}

	if !inplaceChange && scanned == 0 {
		return bs
	}
	b.Write(bs[scanned:len(bs)])
	return []byte(b)
}

func ipliteralEscape(s string) []byte {
	var b bytesp.Slice
	b.WriteByte('[')
	b = isUrlIpLiteralChars.AppendFiltered(b, s[1:len(s)-1])
	b.WriteByte(']')
	return []byte(b)
}

func hostEscape(s string) []byte {
	if len(s) > 4 && s[0] == '[' && s[len(s)-1] == ']' {
		// RFC 3986: IP-literal
		return ipliteralEscape(s)
	}

	return []byte(isUrlRegNameChars.AppendFiltered(nil, s))
}

// htmlReplacer returns s with runes replaced according to replacementTable
// and when badRunes is true, certain bad runes are allowed through unescaped.
func htmlReplacer(s string, replacementTable []HTMLNode, badRunes bool) HTMLNode {
	var b bytesp.Slice
	written := 0
	for i, r := range s {
		if int(r) < len(replacementTable) {
			if repl := replacementTable[r]; len(repl) != 0 {
				b.WriteString(s[written:i])
				b.Write([]byte(repl))
				// Valid as long as replacementTable doesn't
				// include anything above 0x7f.
				written = i + utf8.RuneLen(r)
			}
		} else if badRunes {
			// No-op.
			// IE does not allow these ranges in unquoted attrs.
		} else if 0xfdd0 <= r && r <= 0xfdef || 0xfff0 <= r && r <= 0xffff {
			fmt.Fprintf(&b, "%s&#x%x;", s[written:i], r)
			written = i + utf8.RuneLen(r)
		}
	}
	b.WriteString(s[written:])
	return HTMLNode(b)
}
