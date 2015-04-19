/*
Package htmldef provides some constants definition for hf package.
*/
package htmldef

/**
 * The type of a HTML tag.
 */
type TagType int

const (
	// void elements
	BRTag TagType = iota
	IMGTag
	INPUTTag
	LINKTag
	METATag
	PARAMTag
	
	ATag
	BTag
	BODYTag
	BUTTONTag
	DIVTag
	EMBEDTag
	FOOTERTag
	FORMTag
	H1Tag
	H2Tag
	H3Tag
	H4Tag
	H5Tag
	H6Tag
	HEADTag
	HTMLTag
	LABELTag
	LITag
	NAVTag
	NOSCRIPTTag
	OBJECTTag
	OLTag
	PTag
	PRETag
	SCRIPTTag
	SMALLTag
	SPANTag
	TEXTAREATag
	TITLETag
	ULTag
	
	tagCount
)

/*
TagBytes is a map from a TagType to tag name's bytes.
*/
var TagBytes = [][]byte {
	HTMLTag:     []byte("html"),
	HEADTag:     []byte("head"),
	BODYTag:     []byte("body"),
	METATag:     []byte("meta"),
	TITLETag:    []byte("title"),
	LINKTag:     []byte("link"),
	H1Tag:       []byte("h1"),
	H2Tag:       []byte("h2"),
	H3Tag:       []byte("h3"),
	H4Tag:       []byte("h4"),
	H5Tag:       []byte("h5"),
	H6Tag:       []byte("h6"),
	DIVTag:      []byte("div"),
	PTag:        []byte("p"),
	PRETag:      []byte("pre"),
	SPANTag:     []byte("span"),
	IMGTag:      []byte("img"),
	ATag:        []byte("a"),
	SMALLTag:    []byte("small"),
	BTag:        []byte("b"),
	ULTag:       []byte("ul"),
	LITag:       []byte("li"),
	OLTag:       []byte("ol"),
	FORMTag:     []byte("form"),
	LABELTag:    []byte("label"),
	INPUTTag:    []byte("input"),
	BUTTONTag:   []byte("button"),
	TEXTAREATag: []byte("textarea"),
	SCRIPTTag:   []byte("script"),
	NOSCRIPTTag: []byte("noscript"),
	BRTag:       []byte("br"),
	OBJECTTag:   []byte("object"),
	PARAMTag:    []byte("param"),
	EMBEDTag:    []byte("embed"),
	NAVTag:      []byte("nav"),
	FOOTERTag:   []byte("footer"),
}
