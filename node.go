package vn

import vnd "github.com/mfrachet/go-vdom-wasm/dom"

type Node interface {
	CreateElement()
	GetElement() *vnd.DomElement
	HashCode() string
	ChildrenCount() int
	GetChildren() Children
	GetText() *TextNode
	IsSame(Node) bool
}
