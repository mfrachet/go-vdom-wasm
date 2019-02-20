package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
)

type TextNode struct {
	Value   string
	Element *vnd.DomElement
}

func (textNode *TextNode) SetElement(element vnd.DomElement) {
	textNode.Element = &element
}

func (textNode *TextNode) IsSame(other *TextNode) bool {
	return textNode.Value == other.Value
}
