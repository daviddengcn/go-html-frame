/*
Package htmldef provides some constants definition for hf package.
*/
package htmldef

/**
 * The type of a HTML tag.
 */
type TagType int

const TextType TagType = -1
const (
	// void elements
	AREATag TagType = iota
	BASETag
	BRTag
	COLTag
	HRTag
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
	CAPTIONTag
	COLGROUPTag
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
	HTMLTag
	LABELTag
	LITag
	MAINTag
	MAPTag
	NAVTag
	NOSCRIPTTag
	OBJECTTag
	OLTag
	OPTGROUPTag
	OPTIONTag
	PTag
	PRETag
	RBTag
	RPTag
	RTTag
	RTCTag
	RUBYTag
	SCRIPTTag
	SECTIONTag
	SELECTTag
	SMALLTag
	SPANTag
	TABLETag
	TBODYTag
	TDTag
	TEMPLATETag
	TEXTAREATag
	TFOOTTag
	THTag
	THEADTag
	TITLETag
	TRTag
	ULTag

	tagCount
)

/*
TagBytes is a map from a TagType to tag name
*/
var TagNames = []string{
	AREATag:  "area",
	BASETag:  "base",
	BRTag:    "br",
	COLTag:   "col",
	HRTag:    "hr",
	IMGTag:   "img",
	INPUTTag: "input",
	LINKTag:  "link",
	METATag:  "meta",
	PARAMTag: "param",

	ATag:          "a",
	ADDRESSTag:    "address",
	ARTICLETag:    "article",
	ASIDETag:      "aside",
	BTag:          "b",
	BLOCKQUOTETag: "blockquote",
	BODYTag:       "body",
	BUTTONTag:     "button",
	CAPTIONTag:    "caption",
	COLGROUPTag:   "colgroup",
	DDTag:         "dd",
	DIVTag:        "div",
	DLTag:         "dl",
	DTTag:         "dt",
	EMBEDTag:      "embed",
	FORMTag:       "form",
	FIELDSETTag:   "filedset",
	FOOTERTag:     "footer",
	H1Tag:         "h1",
	H2Tag:         "h2",
	H3Tag:         "h3",
	H4Tag:         "h4",
	H5Tag:         "h5",
	H6Tag:         "h6",
	HEADTag:       "head",
	HEADERTag:     "header",
	HGROUPTag:     "hgroup",
	HTMLTag:       "html",
	LABELTag:      "label",
	LITag:         "li",
	MAINTag:       "main",
	MAPTag:        "map",
	NAVTag:        "nav",
	NOSCRIPTTag:   "noscript",
	OBJECTTag:     "object",
	OLTag:         "ol",
	OPTGROUPTag:   "optgroup",
	OPTIONTag:     "option",
	PTag:          "p",
	PRETag:        "pre",
	RBTag:         "rb",
	RPTag:         "rp",
	RTTag:         "rt",
	RTCTag:        "rtc",
	RUBYTag:       "ruby",
	SCRIPTTag:     "script",
	SECTIONTag:    "section",
	SELECTTag:     "select",
	SMALLTag:      "small",
	SPANTag:       "span",
	TABLETag:      "table",
	TBODYTag:      "tbody",
	TDTag:         "td",
	TEXTAREATag:   "textarea",
	TEMPLATETag:   "template",
	TFOOTTag:      "tfoot",
	THTag:         "th",
	THEADTag:      "thead",
	TITLETag:      "title",
	TRTag:         "tr",
	ULTag:         "ul",
}
