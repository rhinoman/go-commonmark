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

func TestCMarkParser(t *testing.T) {
	parser := commonmark.NewCmarkDocParser()
	if parser == nil {
		t.Error("Parser is nil!")
	}
	parser.ProcessLine("Boo\n")
	parser.ProcessLine("===\n")
	document := parser.Finish()
	if document == nil {
		t.Error("Document is nil!")
	}
	parser.Free()
	htmlText := document.RenderHtml()
	if htmlText != "<h1>Boo</h1>\n" {
		t.Error("Html text is not as expected :(")
	}
	t.Logf("Html Text: %v", htmlText)
	document.DebugPrint()
	document.Free()

	document2 := commonmark.ParseDocument("Foobar\n------")
	htmlText = document2.RenderHtml()
	if htmlText != "<h2>Foobar</h2>\n" {
		t.Error("Html text 2 is not as expected :(")
	}
	t.Logf("Html Text2: %v", htmlText)
	document2.Free()
}
