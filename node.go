package commonmark

/*
#include <stdlib.h>
#include "cmark.h"
*/
import "C"
import "unsafe"

//Maps to a cmark_node_type enum in cmark.h
type NodeType int

const (
	CMARK_NODE_DOCUMENT NodeType = iota
	CMARK_NODE_BLOCK_QUOTE
	CMARK_NODE_LIST
	CMARK_NODE_LIST_ITEM
	CMARK_NODE_CODE_BLOCK
	CMARK_NODE_HTML
	CMARK_NODE_PARAGRAPH
	CMARK_NODE_HEADER
	CMARK_NODE_HRULE
	CMARK_NODE_REFERENCE_DEF

	// Inline
	CMARK_NODE_STRING
	CMARK_NODE_SOFTBREAK
	CMARK_NODE_LINEBREAK
	CMARK_NODE_INLINE_CODE
	CMARK_NODE_INLINE_HTML
	CMARK_NODE_EMPH
	CMARK_NODE_STRONG
	CMARK_NODE_LINK
	CMARK_NODE_IMAGE

	CMARK_NODE_FIRST_BLOCK = CMARK_NODE_DOCUMENT
	CMARK_NODE_LAST_BLOCK = CMARK_NODE_REFERENCE_DEF
	CMARK_NODE_FIRST_INLINE = CMARK_NODE_STRING
	CMARK_NODE_LAST_INLINE = CMARK_NODE_IMAGE
)

//Wraps the cmark_node.
//CommonMark nodes are represented as Trees in memory.
type CMarkNode struct {
	node *C.struct_cmark_node
}

func NewCMarkNode(nt NodeType) *CMarkNode {
	return &CMarkNode{
		node: C.cmark_node_new(C.cmark_node_type(nt)),
	}
}

// Debug print
func (node *CMarkNode) DebugPrint() {
	C.cmark_debug_print(node.node)
}

// Renders the document as HTML.
// Retunrs an HTML string.
func (node *CMarkNode) RenderHtml() string {
	result := C.cmark_render_html(node.node)
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

// Unlink a node from the tree and free it.
func (node *CMarkNode) Destroy() {
	C.cmark_node_destroy(node.node)
}

// Cleanup the Nodelist, including any children
// Once you call Free on this, you can't use it anymore
func (node *CMarkNode) Free() {
	C.cmark_free_nodes(node.node)
}

//Node traversal functions

//Get next node
func (node *CMarkNode) Next() *CMarkNode {
	return &CMarkNode{
		node: C.cmark_node_next(node.node),
	}
}

//Get previous node
func (node *CMarkNode) Previous() *CMarkNode {
	return &CMarkNode{
		node: C.cmark_node_previous(node.node),
	}
}

//Get parent node
func (node *CMarkNode) Parent() *CMarkNode {
	return &CMarkNode{
		node: C.cmark_node_parent(node.node),
	}
}

//Get first child node
func (node *CMarkNode) FirstChild() *CMarkNode {
	return &CMarkNode{
		node: C.cmark_node_first_child(node.node),
	}
}

//Get last child node
func (node *CMarkNode) LastChild() *CMarkNode {
	return &CMarkNode{
		node: C.cmark_node_last_child(node.node),
	}
}

//Accessor functions

//Get the node type
func (node *CMarkNode) GetNodeType() NodeType {
	nt := C.cmark_node_get_type(node.node)
	return NodeType(nt)
}

//Get the node's string content
func (node *CMarkNode) GetStringContent() string {
	cstr := C.cmark_node_get_string_content(node.node)
	return C.GoString(cstr)
}

//Set the node's string content
func (node *CMarkNode) SetStringContent(content string) {
	cstr := C.CString(content)
	defer C.free(unsafe.Pointer(cstr))
	C.cmark_node_set_string_content(node.node, cstr)
}

//Header node funcs

//Get a Header node's level
func (node *CMarkNode) GetHeaderLevel() int {
	level := C.cmark_node_get_header_level(node.node)
	return int(level)
}

//Set a Header node's level (1,2, etc.)
func (node *CMarkNode) SetHeaderLevel(level int) {
	C.cmark_node_set_header_level(node.node, C.int(level))
}

// Tree manipulation functions

//Unlink a node from the tree
func (node *CMarkNode) Unlink() {
	C.cmark_node_unlink(node.node)
}

// InsertBefore can cause a panic quite readily :)
// Hint: Both nodes had better already be in the 'tree'
// Insert a node before another 'sibling' node
func (node *CMarkNode) InsertBefore(sibling *CMarkNode) {
	C.cmark_node_insert_before(node.node, sibling.node)
}

// InsertAfter can cause a panic quite readily :)
// Hint: Both nodes had better already be in the 'tree'
//Insert a node after another 'sibling' node
func (node *CMarkNode) InsertAfter(sibling *CMarkNode) {
	C.cmark_node_insert_after(node.node, sibling.node)
}

//Prepend a child node
func (node *CMarkNode) PrependChild(child *CMarkNode) {
	C.cmark_node_prepend_child(node.node, child.node)
}

//Append a child node
func (node *CMarkNode) AppendChild(child *CMarkNode) {
	C.cmark_node_append_child(node.node, child.node)
}
