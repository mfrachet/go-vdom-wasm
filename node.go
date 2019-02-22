package vn

import vnd "github.com/mfrachet/go-vdom-wasm/dom"

type Node interface {
	GetElement() *vnd.DomNode
	GetTagName() string
	GetAttrs() *Attrs
	GetKey() *string
	HasElement() bool
	SetElement(vnd.DomNode)
	HashCode() string
	ChildrenCount() int
	ChildAt(int) Node
	GetText() *TextNode
	GetChildren() Children
	IsSame(Node) bool
}
