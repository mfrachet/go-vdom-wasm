package vn

import vnd "github.com/mfrachet/go-vdom-wasm/dom"

type Node interface {
	createElement()
	getElement() *vnd.DomElement
	hashCode() string
	childrenCount() int
	getChildren() Children
	getText() *TextNode
	isSame(Node) bool
}
