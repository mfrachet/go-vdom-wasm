package vn

import (
	"syscall/js"
)

func CreateElement(node interface{}) *js.Value {
	switch node.(type) {
	case *TextNode:
		vnode := node.(*TextNode)
		vnode.createElement()

		return vnode.Element

	default:
		vnode := node.(*Vnode)
		vnode.createElement()

		return vnode.Element
	}
}
