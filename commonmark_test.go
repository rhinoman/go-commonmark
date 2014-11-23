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
	document2.DebugPrint()
	if htmlText != "<h2>Foobar</h2>\n" {
		t.Error("Html text 2 is not as expected :(")
	}
	t.Logf("Html Text2: %v", htmlText)
	document2.Free()
}

func TestCMarkNodeOps(t *testing.T) {
	root := commonmark.NewCMarkNode(commonmark.CMARK_NODE_DOCUMENT)
	if root == nil {
		t.Error("Root is nil!")
	}
	if root.GetNodeType() != commonmark.CMARK_NODE_DOCUMENT {
		t.Error("Root is wrong type!")
	}
	header1 := commonmark.NewCMarkNode(commonmark.CMARK_NODE_HEADER)
	if header1.GetNodeType() != commonmark.CMARK_NODE_HEADER {
		t.Error("header1 is wrong type!")
	}
	header1.SetHeaderLevel(1)
	header1str := commonmark.NewCMarkNode(commonmark.CMARK_NODE_STRING)
	header1str.SetStringContent("I'm the main header!")
	if header1str.GetStringContent() != "I'm the main header!" {
		t.Error("header1str content is wrong!")
	}
	header1.AppendChild(header1str)
	header2 := commonmark.NewCMarkNode(commonmark.CMARK_NODE_HEADER)
	header2str := commonmark.NewCMarkNode(commonmark.CMARK_NODE_STRING)
	header2str.SetStringContent("Another header!")
	header2.AppendChild(header2str)
	header2.SetHeaderLevel(2)
	root.PrependChild(header1)
	root.AppendChild(header2)
	root.DebugPrint()
	htmlStr := root.RenderHtml()
	if htmlStr != "<h1>I'm the main header!</h1>\n<h2>Another header!</h2>\n" {
		t.Error("htmlStr is wrong!")
	}
	t.Logf("Html Text: %v", htmlStr)
	//Rearrange...
	header1.InsertBefore(header2)
	root.DebugPrint()
	htmlStr = root.RenderHtml()
	if htmlStr != "<h2>Another header!</h2>\n<h1>I'm the main header!</h1>\n" {
		t.Error("htmlStr is wrong!")
	}
	t.Logf("Html Text: %v", htmlStr)
	//removing something
	header2.Unlink()
	root.DebugPrint()
	htmlStr = root.RenderHtml()
	if htmlStr != "<h1>I'm the main header!</h1>\n" {
		t.Error("htmlStr is wrong!")
	}
	header2.Destroy()
	root.Free()
}
