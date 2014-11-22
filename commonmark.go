package commonmark

// #cgo LDFLAGS: -lcmark
// #include "cmark.h"
import "C"
import "strings"

//Converts Markdo--, er, CommonMark text to Html
//mdtext -- input CommonMark text
//returns HTML string
func Md2Html(mdtext string) string {
	//The call to cmark will barf if the input string doesn't end with a newline
	if !strings.HasSuffix(mdtext, "\n") {
		mdtext += "\n"
	}
	mdCstr := C.CString(mdtext)
	strLen := C.int(len(mdtext))
	htmlString := C.cmark_markdown_to_html(mdCstr, strLen)
	return C.GoString(htmlString)
}
