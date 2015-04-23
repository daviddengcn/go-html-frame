package hf

import (
	"bytes"
	"strconv"
)

func isSpaceCharacters(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\f' || b == '\r'
}

func startWithSpace(txt HTMLBytes) bool {
	if len(txt) == 0 {
		return false
	}

	return !isSpaceCharacters(txt[0])
}

func findStrInArr(s HTMLBytes, arr []HTMLBytes) int {
	for i, el := range arr {
		if bytes.Equal(el, s) {
			return i
		}
	}

	return -1
}

func intSliceToBytes(ints []int) HTMLBytes {
	var b []byte
	for idx, i := range ints {
		b = strconv.AppendInt(b, int64(i), 10)
		if idx > 0 {
			b = append(b, ',')
		}
	}
	return HTMLBytes(b)
}

func itoaBytes(i int) HTMLBytes {
	return HTMLBytes(strconv.AppendInt(nil, int64(i), 10))
}