//Package commonmark provides a Go wrapper for the CommonMark C Library
package commonmark

/*
#cgo LDFLAGS: -lcmark
#include <stdlib.h>
#include "cmark.h"
*/
import "C"
import (
	"strings"
	"unsafe"
)

// Converts Markdo--, er, CommonMark text to Html.
// Parameter mdtext contains CommonMark text.
// The return value is the HTML string
func Md2Html(mdtext string) string {
	//The call to cmark will barf if the input string doesn't end with a newline
	if !strings.HasSuffix(mdtext, "\n") {
		mdtext += "\n"
	}
	mdCstr := C.CString(mdtext)
	strLen := C.int(len(mdtext))
	defer C.free(unsafe.Pointer(mdCstr))
	htmlString := C.cmark_markdown_to_html(mdCstr, strLen)
	defer C.free(unsafe.Pointer(htmlString))
	return C.GoString(htmlString)
}

//Wraps the cmark_doc_parser
type CMarkParser struct {
	parser *C.struct_cmark_doc_parser
}

// Retruns a new CMark Parser.
// You must call Free() on this thing when you're done with it!
// Please.
func NewCmarkDocParser() *CMarkParser {
	return &CMarkParser{
		parser: C.cmark_new_doc_parser(),
	}
}

// Process a line
func (cmp *CMarkParser) ProcessLine(line string) {
	s := len(line)
	cLine := C.CString(line)
	defer C.free(unsafe.Pointer(cLine))
	C.process_line(cmp.parser, cLine, C.size_t(s))
}

// Finish parsing and generate a document
// You must call Free() on the document when you're done with it!
func (cmp *CMarkParser) Finish() *CMarkNode {
	return &CMarkNode{
		node: C.cmark_finish(cmp.parser),
	}
}

// Cleanup the parser
// Once you call Free on this, you can't use it anymore
func (cmp *CMarkParser) Free() {
	C.cmark_free_doc_parser(cmp.parser)
}

// Generates a document directly from a string
func ParseDocument(buffer string) *CMarkNode {
	if !strings.HasSuffix(buffer, "\n") {
		buffer += "\n"
	}
	Cstr := C.CString(buffer)
	Clen := C.size_t(len(buffer))
	defer C.free(unsafe.Pointer(Cstr))
	return &CMarkNode{
		node: C.cmark_parse_document(Cstr, Clen),
	}
}
