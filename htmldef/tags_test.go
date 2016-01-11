package htmldef

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestTagNames(t *testing.T) {
	assert.Equal(t, "Length of TagNames", len(TagNames), int(tagCount))

	last := ""
	for i, name := range TagNames {
		if len(name) == 0 {
			if len(last) != 0 {
				t.Errorf("Bytes for tag after %s is undefined.", last)
			} else {
				t.Errorf("Bytes for tag of value %d is undefined.", i)
			}
		}

		last = name
	}
}
