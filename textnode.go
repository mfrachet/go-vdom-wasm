package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
)

type TextNode struct {
	Value   string
	Element *vnd.DomNode
}

func NewTextnode(value string) *TextNode {
	return &TextNode{value, nil}
}

func (textNode *TextNode) GetElement() *vnd.DomNode {
	return textNode.Element
}

func (textNode *TextNode) GetValue() string {
	return textNode.Value
}

func (textNode *TextNode) SetElement(element vnd.DomNode) {
	textNode.Element = &element
}

func (textNode *TextNode) IsSame(other *TextNode) bool {
	return textNode.GetValue() == other.GetValue()

}
