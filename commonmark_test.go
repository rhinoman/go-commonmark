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
	parser := commonmark.NewCmarkParser()
	if parser == nil {
		t.Error("Parser is nil!")
	}
	parser.Feed("Boo\n")
	parser.Feed("===\n")
	document := parser.Finish()
	if document == nil {
		t.Error("Document is nil!")
	}
	//Call it twice to make sure it doesn't crash :)
	parser.Free()
	parser.Free()
	htmlText := document.RenderHtml()
	if htmlText != "<h1>Boo</h1>\n" {
		t.Error("Html text is not as expected :(")
	}
	t.Logf("Html Text: %v", htmlText)
	document.RenderAst()
	document.Free()

	document2 := commonmark.ParseDocument("Foobar\n------")
	htmlText = document2.RenderHtml()
	document2.RenderAst()
	if htmlText != "<h2>Foobar</h2>\n" {
		t.Error("Html text 2 is not as expected :(")
	}
	t.Logf("Html Text2: %v", htmlText)
	document2.Free()
	document2.Free()
}

func TestParseFile(t *testing.T) {
	node, err := commonmark.ParseFile("test_data/test_file.md")
	if err != nil {
		t.Error(err)
	}
	if node == nil {
		t.Error(err)
	}
	htmlText := node.RenderHtml()
	if htmlText != "<h1>Test File</h1>\n<h2>Description</h2>\n<p>This is just a test file.</p>\n" {
		t.Error("Html text is not as expected :(")
	}
	t.Logf("Html Text: %v", htmlText)
	node.Free()
	//try to parse a non-existent file
	eNode, err := commonmark.ParseFile("notafile.md")
	if err == nil {
		t.Errorf("Should have been an error!")
	}
	t.Logf("error string: %v", err.Error())
	if eNode != nil {
		t.Errorf("Node should be nil!")
	}
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
	if header1.SetLiteral("boo") != false {
		t.Error("SetLiteral should return false for header node")
	}
	header1str := commonmark.NewCMarkNode(commonmark.CMARK_NODE_TEXT)
	header1str.SetLiteral("I'm the main header!")
	if header1str.GetLiteral() != "I'm the main header!" {
		t.Error("header1str content is wrong!")
	}
	header1.AppendChild(header1str)
	header2 := commonmark.NewCMarkNode(commonmark.CMARK_NODE_HEADER)
	header2str := commonmark.NewCMarkNode(commonmark.CMARK_NODE_TEXT)
	if header2str.SetLiteral("Another header!") == false {
		t.Error("SetLiteral returned false for valid input")
	}
	header2.AppendChild(header2str)
	header2.SetHeaderLevel(2)
	if root.PrependChild(header1) == false {
		t.Error("Couldn't prepend header to root")
	}
	root.AppendChild(header2)
	t.Logf("\nAST: %v", root.RenderAst())
	htmlStr := root.RenderHtml()
	if htmlStr != "<h1>I'm the main header!</h1>\n<h2>Another header!</h2>\n" {
		t.Error("htmlStr is wrong!")
	}
	t.Logf("Html Text: %v", htmlStr)
	//Rearrange...
	header1.InsertBefore(header2)
	t.Logf("\nAST: %v", root.RenderAst())
	htmlStr = root.RenderHtml()
	if htmlStr != "<h2>Another header!</h2>\n<h1>I'm the main header!</h1>\n" {
		t.Error("htmlStr is wrong!")
	}
	t.Logf("Html Text: %v", htmlStr)
	//removing something
	header2.Unlink()
	t.Logf("\nAST: %v", root.RenderAst())
	htmlStr = root.RenderHtml()
	if htmlStr != "<h1>I'm the main header!</h1>\n" {
		t.Error("htmlStr is wrong!")
	}
	header2.Free()
	root.Free()
}

func TestCMarkLists(t *testing.T) {
	root := commonmark.NewCMarkNode(commonmark.CMARK_NODE_DOCUMENT)
	list := commonmark.NewCMarkNode(commonmark.CMARK_NODE_LIST)
	list.SetListType(commonmark.CMARK_ORDERED_LIST)
	listItem1 := commonmark.NewCMarkNode(commonmark.CMARK_NODE_LIST_ITEM)
	listItem2 := commonmark.NewCMarkNode(commonmark.CMARK_NODE_LIST_ITEM)
	li1para := commonmark.NewCMarkNode(commonmark.CMARK_NODE_PARAGRAPH)
	li1str := commonmark.NewCMarkNode(commonmark.CMARK_NODE_TEXT)
	li1str.SetLiteral("List Item 1")
	li1para.AppendChild(li1str)
	if listItem1.AppendChild(li1para) == false {
		t.Error("Couldn't append paragraph to list item")
	}
	list.AppendChild(listItem1)
	list.AppendChild(listItem2)
	list.SetListTight(true)
	root.AppendChild(list)
	t.Logf("\nAST: %v", root.RenderAst())
	htmlString := root.RenderHtml()
	if htmlString != "<ol>\n<li>List Item 1</li>\n<li></li>\n</ol>\n" {
		t.Error("htmlString is wrong!")
	}
	t.Logf("\nHtmlString: \n%v", htmlString)
	t.Logf("\nList start: %v", list.GetListStart())
	t.Logf("\nList tight: %v", list.GetListTight())
	root.Free()
}

func TestCMarkCodeBlocks(t *testing.T) {
	root := commonmark.NewCMarkNode(commonmark.CMARK_NODE_DOCUMENT)
	cb := commonmark.NewCMarkNode(commonmark.CMARK_NODE_CODE_BLOCK)
	cb.SetLiteral("int main(){\n return 0;\n }")
	cb.SetFenceInfo("c")
	if cb.GetFenceInfo() != "c" {
		t.Error("Fence info isn't c")
	}
	if cb.GetLiteral() != "int main(){\n return 0;\n }" {
		t.Error("Code has changed somehow")
	}
	if root.AppendChild(cb) == false {
		t.Error("Couldn't append code block to document")
	}
	t.Logf("\nAST: %v", root.RenderAst())
	htmlString := root.RenderHtml()
	t.Logf("\nHtml String: %v\n", htmlString)
	if htmlString != "<pre><code class=\"language-c\">int main(){\n return 0;\n }</code></pre>\n" {
		t.Error("htmlString isn't right!")
	}
	root.Free()
}

func TestCMarkUrls(t *testing.T) {
	root := commonmark.NewCMarkNode(commonmark.CMARK_NODE_DOCUMENT)
	para := commonmark.NewCMarkNode(commonmark.CMARK_NODE_PARAGRAPH)
	link := commonmark.NewCMarkNode(commonmark.CMARK_NODE_LINK)
	root.AppendChild(para)
	if para.AppendChild(link) == false {
		t.Error("Couldn't append link node to paragraph!")
	}
	if link.SetUrl("http://duckduckgo.com") == false {
		t.Error("Couldn't set URL!!!")
	}
	if link.GetUrl() != "http://duckduckgo.com" {
		t.Error("Url doesn't match")
	}
	t.Logf("\nAST: %v", root.RenderAst())
	htmlString := root.RenderHtml()
	t.Logf("\nHtml String: %v\n", htmlString)
	if htmlString != "<p><a href=\"http://duckduckgo.com\"></a></p>\n" {
		t.Error("htmlString isn't right!")
	}
	root.Free()
}
