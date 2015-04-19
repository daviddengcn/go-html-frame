/*
Package htmldef provides some constants definition for hf package.
*/
package htmldef

/**
 * The type of a HTML tag.
 */
type TagType int

const (
	TextType TagType = -1
	
	// void elements
	BRTag TagType = iota
	IMGTag
	INPUTTag
	LINKTag
	METATag
	PARAMTag
	
	ATag
	ADDRESSTag
	ARTICLETag
	ASIDETag
	BTag
	BLOCKQUOTETag
	BODYTag
	BUTTONTag
	DDTag
	DIVTag
	DLTag
	DTTag
	EMBEDTag
	FIELDSETTag
	FOOTERTag
	FORMTag
	H1Tag
	H2Tag
	H3Tag
	H4Tag
	H5Tag
	H6Tag
	HEADTag
	HEADERTag
	HGROUPTag
	HRTag
	HTMLTag
	LABELTag
	LITag
	MAINTag
	NAVTag
	NOSCRIPTTag
	OBJECTTag
	OLTag
	PTag
	PRETag
	RBTag
	RPTag
	RTTag
	RTCTag
	RUBYTag
	SCRIPTTag
	SECTIONTag
	SMALLTag
	SPANTag
	TABLETag
	TEMPLATETag
	TEXTAREATag
	TITLETag
	ULTag
	
	tagCount
)

/*
TagBytes is a map from a TagType to tag name's bytes.
*/
var TagBytes = [][]byte {
	BRTag:       []byte("br"),
	IMGTag:      []byte("img"),
	INPUTTag:    []byte("input"),
	LINKTag:     []byte("link"),
	METATag:     []byte("meta"),
	PARAMTag:    []byte("param"),
	
	ATag:        []byte("a"),
	ADDRESSTag:  []byte("address"),
	ARTICLETag:  []byte("article"),
	ASIDETag:    []byte("aside"),
	BTag:        []byte("b"),
	BLOCKQUOTETag: []byte("blockquote"),
	BODYTag:     []byte("body"),
	BUTTONTag:   []byte("button"),
	DDTag:       []byte("dd"),
	DIVTag:      []byte("div"),
	DLTag:       []byte("dl"),
	DTTag:       []byte("dt"),
	EMBEDTag:    []byte("embed"),
	FORMTag:     []byte("form"),
	FIELDSETTag: []byte("filedset"),
	FOOTERTag:   []byte("footer"),
	H1Tag:       []byte("h1"),
	H2Tag:       []byte("h2"),
	H3Tag:       []byte("h3"),
	H4Tag:       []byte("h4"),
	H5Tag:       []byte("h5"),
	H6Tag:       []byte("h6"),
	HEADTag:     []byte("head"),
	HEADERTag:   []byte("header"),
	HGROUPTag:   []byte("hgroup"),
	HRTag:       []byte("hr"),
	HTMLTag:     []byte("html"),
	LABELTag:    []byte("label"),
	LITag:       []byte("li"),
	MAINTag:     []byte("main"),
	NAVTag:      []byte("nav"),
	NOSCRIPTTag: []byte("noscript"),
	OBJECTTag:   []byte("object"),
	OLTag:       []byte("ol"),
	PTag:        []byte("p"),
	PRETag:      []byte("pre"),
	RBTag:       []byte("rb"),
	RPTag:       []byte("rp"),
	RTTag:       []byte("rt"),
	RTCTag:      []byte("rtc"),
	RUBYTag:     []byte("ruby"),
	SCRIPTTag:   []byte("script"),
	SECTIONTag:  []byte("section"),
	SMALLTag:    []byte("small"),
	SPANTag:     []byte("span"),
	TABLETag:    []byte("table"),
	TEXTAREATag: []byte("textarea"),
	TEMPLATETag: []byte("template"),
	TITLETag:    []byte("title"),
	ULTag:       []byte("ul"),
}
