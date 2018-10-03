package vn

import "syscall/js"

type Node interface {
	createElement()
	getElement() *js.Value
	hashCode() string
	childrenCount() int
	getChildren() Children
	getText() *TextNode
	isSame(Node) bool
}
