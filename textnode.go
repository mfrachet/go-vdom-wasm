package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

type TextNode struct {
	Value   string
	Element *vnd.DomElement
}

func (textNode *TextNode) MakeDomNode(document vnd.DomNode) *vnd.DomElement {
	if vnh.IsNil(textNode.Element) {
		domNode := document.CreateTextNode(textNode.Value)
		textNode.Element = &domNode
	}

	return textNode.Element
}

func (textNode *TextNode) GetElement() *vnd.DomElement {
	return textNode.Element
}

func (textNode *TextNode) SetElement(element vnd.DomElement) {
	textNode.Element = &element
}

func (textNode *TextNode) HashCode() string {
	return textNode.Value
}

func (textNode *TextNode) IsSame(other *TextNode) bool {
	return textNode.HashCode() == other.HashCode()
}

func (textNode *TextNode) ChildrenCount() int {
	return -1
}

func (textNode *TextNode) ChildAt(index int) Node {
	return nil
}

func (textNode *TextNode) GetText() *TextNode {
	return textNode
}
