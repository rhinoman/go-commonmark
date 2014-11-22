package commonmark

// #cgo LDFLAGS: -lcmark
// #include "cmark.h"
import "C"
import "strings"

func Md2Html(mdtext string) string {
	//cmark will barf if the input string doesn't end with a newline
	if !strings.HasSuffix(mdtext, "\n") {
		mdtext += "\n"
	}
	mdCstr := C.CString(mdtext)
	strLen := C.int(len(mdtext))
	htmlString := C.cmark_markdown_to_html(mdCstr, strLen)
	return C.GoString(htmlString)
}
