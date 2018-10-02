package vn

import "syscall/js"

type Node interface {
	createElement()
	getElement() *js.Value
	hashCode() string
	childrenCount() int
	getChildren() interface{}
	isSame(Node) bool
}
