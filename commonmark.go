//Package commonmark provides a Go wrapper for the CommonMark C Library
package commonmark

/*
#cgo LDFLAGS: -lcmark
#include <stdio.h>
#include <stdlib.h>
#include "cmark.h"
*/
import "C"
import (
	"errors"
	"runtime"
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
	parser *C.struct_cmark_parser
}

// Retruns a new CMark Parser.
// You must call Free() on this thing when you're done with it!
// Please.
func NewCmarkParser() *CMarkParser {
	p := &CMarkParser{
		parser: C.cmark_parser_new(),
	}
	runtime.SetFinalizer(p, (*CMarkParser).Free)
	return p
}

// Process a line
func (cmp *CMarkParser) Push(line string) {
	s := len(line)
	cLine := C.CString(line)
	defer C.free(unsafe.Pointer(cLine))
	C.cmark_parser_push(cmp.parser, cLine, C.size_t(s))
}

// Finish parsing and generate a document
// You must call Free() on the document when you're done with it!
func (cmp *CMarkParser) Finish() *CMarkNode {
	n := &CMarkNode{
		node: C.cmark_parser_finish(cmp.parser),
	}
	runtime.SetFinalizer(n, (*CMarkNode).Free)
	return n
}

// Cleanup the parser
// Once you call Free on this, you can't use it anymore
func (cmp *CMarkParser) Free() {
	if cmp.parser != nil {
		C.cmark_parser_free(cmp.parser)
	}
	cmp.parser = nil
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

// Parses a file and returns a CMarkNode
// Returns an error if the file can't be opened
func ParseFile(filename string) (*CMarkNode, error) {
	fname := C.CString(filename)
	access := C.CString("r")
	defer C.free(unsafe.Pointer(fname))
	defer C.free(unsafe.Pointer(access))
	file := C.fopen(fname, access)
	if file == nil {
		return nil, errors.New("Unable to open file with name: " + filename)
	}
	defer C.fclose(file)
	n := &CMarkNode{
		node: C.cmark_parse_file(file),
	}
	runtime.SetFinalizer(n, (*CMarkNode).Free)
	return n, nil
}
