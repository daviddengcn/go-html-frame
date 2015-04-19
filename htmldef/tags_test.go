package htmldef

import (
	"testing"
)

func TestTagBytes(t *testing.T) {
	if len(TagBytes) != int(tagCount) {
		t.Errorf("Length of TagBytes is expected to be %d but got %d", tagCount, len(TagBytes))
	}
	
	last := []byte(nil)
	for i, bs := range TagBytes {
		if len(bs) == 0 {
			if len(last) != 0 {
				t.Errorf("Bytes for tag after %s is undefined.", string(last))
			} else {
				t.Errorf("Bytes for tag of value %d is undefined.", i)
			}
		}
		
		last = bs
	}
}