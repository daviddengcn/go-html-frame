package hf

import (
	"strconv"

	"github.com/golangplus/bytes"
)

func isSpaceCharacters(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\f' || b == '\r'
}

func startWithSpace(txt HTMLNode) bool {
	if len(txt) == 0 {
		return false
	}

	return isSpaceCharacters(txt[0])
}

type HTMLNodeSet []HTMLNode

func (set *HTMLNodeSet) Put(s HTMLNode) {
	for _, el := range *set {
		if el == s {
			return
		}
	}

	*set = append(*set, s)
}

func (set *HTMLNodeSet) Del(s HTMLNode) {
	for i, el := range *set {
		if el == s {
			*set = append((*set)[:i], (*set)[i+1:]...)
			return
		}
	}
}

func intSliceToBytes(ints []int) HTMLNode {
	var b []byte
	for idx, i := range ints {
		b = strconv.AppendInt(b, int64(i), 10)
		if idx > 0 {
			b = append(b, ',')
		}
	}
	return HTMLNode(b)
}

func itoaBytes(i int) HTMLNode {
	return HTMLNode(strconv.AppendInt(nil, int64(i), 10))
}

type filterTableArr [256]bool

func (arr filterTableArr) String() string {
	var b bytesp.ByteSlice
	for c, bl := range arr {
		if bl {
			b.WriteByte(byte(c))
		}
	}
	return strconv.Quote(string(b))
}

func filterTableArrFromString(s string) (arr filterTableArr) {
	arr.SetByString(s)
	return
}

// Returns self for chaining grammar
func (arr *filterTableArr) UnionEqual(other filterTableArr) *filterTableArr {
	for i, el := range other {
		if el {
			arr[i] = true
		}
	}

	return arr
}

func (allowed *filterTableArr) AppendFiltered(b bytesp.Slice, s string) bytesp.Slice {
	scanned := 0
	for i, r := range s {
		if !allowed[r] {
			if i > scanned {
				b.WriteString(s[scanned:i])
			}
			scanned = i + 1
		}
	}

	b.WriteString(s[scanned:len(s)])

	return b
}

func (allowed *filterTableArr) AppendPctEncode(b bytesp.Slice, s string) bytesp.Slice {
	scanned := 0
	for i, n := 0, len(s); i < n; i++ {
		r := s[i]
		if !allowed[r] {
			if i > scanned {
				b.WriteString(s[scanned:i])
			}
			b.WriteByte('%')
			b.WriteByte(dec2hex[r/0x10])
			b.WriteByte(dec2hex[r%0x10])

			scanned = i + 1
		}
	}

	b.WriteString(s[scanned:len(s)])

	return b
}

func (arr *filterTableArr) SetByString(s string) {
	for _, c := range s {
		if c < rune(len(arr)) {
			arr[c] = true
		}
	}
}

func (arr *filterTableArr) SetRange(mn, mx byte) {
	for i := int(mn); i <= int(mx); i++ {
		arr[i] = true
	}
}
