//Package commonmark provides a Go wrapper for the CommonMark C Library
package commonmark

/*
#cgo LDFLAGS: -lcmark
#include <stdlib.h>
#include "cmark.h"
*/
import "C"
import "strings"
import "unsafe"

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

type CMarkParser struct {
	parser *C.struct_cmark_doc_parser
}

type CMarkDocument struct {
	document *C.struct_cmark_node
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
func (cmp *CMarkParser) ProcessLine(buffer string) {
	s := len(buffer)
	C.process_line(cmp.parser, C.CString(buffer), C.size_t(s))
}

// Finish parsing and generate a document
// You must call Free() on the document when you're done with it!
func (cmp *CMarkParser) Finish() *CMarkDocument {
	return &CMarkDocument{
		document: C.cmark_finish(cmp.parser),
	}
}

// Cleanup the parser
// Once you call Free on this, you can't use it anymore
func (cmp *CMarkParser) Free() {
	C.cmark_free_doc_parser(cmp.parser)
}

// Renders the document as HTML.
// Retunrs an HTML string.
func (doc *CMarkDocument) RenderHtml() string {
	result := C.cmark_render_html(doc.document)
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

// Cleanup the document
// Once you call Free on this, you can't use it anymore
func (doc *CMarkDocument) Free() {
	C.cmark_free_nodes(doc.document)
}
