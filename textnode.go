package vn

import (
	vn_dom "github.com/mfrachet/go-vdom-wasm/dom"
)

type TextNode struct {
	Value   string
	Element *vn_dom.DomElement
}

func (textNode *TextNode) createElement() {
	if textNode.Element == nil {
		document := vn_dom.GetDocument()
		domNode := document.CreateTextNode(textNode.Value)
		textNode.Element = &domNode
	}
}

func (textNode *TextNode) getElement() *vn_dom.DomElement {
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
