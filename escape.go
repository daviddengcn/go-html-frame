package hf

import (
	"fmt"
	"html/template"
	"unicode/utf8"

	"github.com/daviddengcn/go-villa"
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
var htmlReplacementTable = []HTMLBytes{
	// http://www.w3.org/TR/html5/syntax.html#attribute-value-(unquoted)-state
	// U+0000 NULL Parse error. Append a U+FFFD REPLACEMENT
	// CHARACTER character to the current attribute's value.
	// "
	// and similarly
	// http://www.w3.org/TR/html5/syntax.html#before-attribute-value-state
	0:    HTMLBytes("\uFFFD"),
	0xA0: NBSP,
	'"':  HTMLBytes("&quot;"),
	'&':  AMP,
	'<':  LT,
	'>':  GT,
}

// htmlReplacementTable contains the runes that need to be escaped
// inside a quoted attribute value or in a text node.
var attrReplacementTable = []HTMLBytes{
	// http://www.w3.org/TR/html5/syntax.html#attribute-value-(unquoted)-state
	// U+0000 NULL Parse error. Append a U+FFFD REPLACEMENT
	// CHARACTER character to the current attribute's value.
	// "
	// and similarly
	// http://www.w3.org/TR/html5/syntax.html#before-attribute-value-state
	0:    HTMLBytes("\uFFFD"),
	0xA0: NBSP,
	'"':  HTMLBytes("&quot;"),
	'&':  AMP,
}

// attrEscaper escapes for inclusion in quoted attribute values.
func attrEscaper(attr string) HTMLBytes {
	return htmlReplacer(attr, attrReplacementTable, true)
}

// htmlEscaper escapes for inclusion in HTML text.
func htmlEscaper(text string) HTMLBytes {
	return htmlReplacer(text, htmlReplacementTable, true)
}

// htmlReplacer returns s with runes replaced according to replacementTable
// and when badRunes is true, certain bad runes are allowed through unescaped.
func htmlReplacer(s string, replacementTable []HTMLBytes, badRunes bool) HTMLBytes {
	var b villa.ByteSlice
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
	return HTMLBytes(b)
}
