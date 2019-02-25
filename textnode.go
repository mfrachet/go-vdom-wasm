package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
)

type TextNode struct {
	value   string
	element *vnd.DomNode
}

func NewTextnode(value string) *TextNode {
	return &TextNode{value, nil}
}

func (textNode *TextNode) GetElement() *vnd.DomNode {
	return textNode.element
}

func (textNode *TextNode) GetValue() string {
	return textNode.value
}

func (textNode *TextNode) SetElement(element vnd.DomNode) {
	textNode.element = &element
}

func (textNode *TextNode) IsSame(other *TextNode) bool {
	return textNode.GetValue() == other.GetValue()

}
