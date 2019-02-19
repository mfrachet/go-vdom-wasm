package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

type TextNode struct {
	Value   string
	Element *vnd.DomElement
}

func (textNode *TextNode) createElement() {
	if vnh.IsNil(textNode.Element) {
		document := vnd.GetDocument()
		domNode := document.CreateTextNode(textNode.Value)
		textNode.Element = &domNode
	}
}

func (textNode *TextNode) getElement() *vnd.DomElement {
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

func (textNode *TextNode) getChildren() Children {
	return nil
}

func (textNode *TextNode) getText() *TextNode {
	return textNode
}
