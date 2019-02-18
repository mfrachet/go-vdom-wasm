package vn

import vn_dom "github.com/mfrachet/go-vdom-wasm/dom"

type Node interface {
	createElement()
	getElement() *vn_dom.DomElement
	hashCode() string
	childrenCount() int
	getChildren() Children
	getText() *TextNode
	isSame(Node) bool
}
