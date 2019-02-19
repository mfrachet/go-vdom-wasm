package vn

import vnd "github.com/mfrachet/go-vdom-wasm/dom"

type Node interface {
	MakeDomNode(vnd.DomNode) *vnd.DomElement
	GetElement() *vnd.DomElement
	SetElement(vnd.DomElement)
	HashCode() string
	ChildrenCount() int
	ChildAt(int) Node
	GetText() *TextNode
	IsSame(Node) bool
}
