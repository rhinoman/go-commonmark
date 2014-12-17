package commonmark

/*
#include <stdlib.h>
#include "cmark.h"
*/
import "C"
import (
	"runtime"
)

type NodeEvent int

const (
	CMARK_EVENT_DONE NodeEvent = iota
	CMAR_EVENT_ENTER
	CMARK_EVENT_EXIT
)

//Wraps a cmark_iter
type CMarkIter struct {
	iter *C.cmark_iter
}

//Creates a new iterator starting with the given node.
func NewCMarkIter(node *CMarkNode) *CMarkIter {
	iter := &CMarkIter{
		iter: C.cmark_iter_new(node.node),
	}
	runtime.SetFinalizer(iter, (*CMarkIter).Free)
	return iter
}

//Returns the event type for the next node
func (iter *CMarkIter) Next() NodeEvent {
	ne := C.cmark_iter_next(iter.iter)
	return NodeEvent(ne)
}

//Returns the next node in the sequence
func (iter *CMarkIter) GetNode() *CMarkNode {
	return &CMarkNode{
		node: C.cmark_iter_get_node(iter.iter),
	}

}

//Frees an iterator
func (iter *CMarkIter) Free() {
	if iter.iter != nil {
		C.cmark_iter_free(iter.iter)
	}
	iter.iter = nil
}
