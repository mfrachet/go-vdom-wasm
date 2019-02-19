package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

type TextNode struct {
	Value   string
	Element *vnd.DomElement
}

func (textNode *TextNode) MakeDomNode(document vnd.DomNode) {
	if vnh.IsNil(textNode.Element) {
		domNode := document.CreateTextNode(textNode.Value)
		textNode.Element = &domNode
	}
}

func (textNode *TextNode) GetElement() *vnd.DomElement {
	return textNode.Element
}

func (textNode *TextNode) HashCode() string {
	return textNode.Value
}

func (textNode *TextNode) IsSame(other Node) bool {
	return textNode.HashCode() == other.HashCode()
}

func (textNode *TextNode) ChildrenCount() int {
	return -1
}

func (textNode *TextNode) GetChildren() Children {
	return nil
}

func (textNode *TextNode) GetText() *TextNode {
	return textNode
}
