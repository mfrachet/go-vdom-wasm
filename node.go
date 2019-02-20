package vn

import vnd "github.com/mfrachet/go-vdom-wasm/dom"

type Node interface {
	GetElement() *vnd.DomElement
	GetTagName() string
	GetAttrs() *Attrs
	SetElement(vnd.DomElement)
	HashCode() string
	ChildrenCount() int
	ChildAt(int) Node
	GetText() *TextNode
	GetChildren() Children
	IsSame(Node) bool
}
