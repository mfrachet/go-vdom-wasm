package vn

import (
	"syscall/js"
)

type TextNode struct {
	Value   string
	Element *js.Value
}

func (textNode *TextNode) createElement() {
	if textNode.Element == nil {
		document := js.Global().Get("document")
		domNode := document.Call("createTextNode", textNode.Value)
		textNode.Element = &domNode
	}
}

func (textNode *TextNode) getElement() *js.Value {
	return textNode.Element
}

func (textNode *TextNode) hashCode() string {
	return textNode.Value
}

func (textNode *TextNode) isSame(other Node) bool {
	return textNode.hashCode() == other.hashCode()
}

func (textNode *TextNode) childrenCount() int {
	return -1
}

func (textNode *TextNode) getChildren() interface{} {
	return textNode.Value
}
