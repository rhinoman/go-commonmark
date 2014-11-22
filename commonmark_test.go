package commonmark_test

import (
	"github.com/rhinoman/go-commonmark"
	"testing"
)

func TestMd2Html(t *testing.T) {
	htmlText := commonmark.Md2Html("Boo\n===")
	if htmlText != "<h1>Boo</h1>\n" {
		t.Errorf("Html text is not as expected :(")
	}
	t.Logf("Html Text: %v", htmlText)
}
