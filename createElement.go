package vn

import (
	"syscall/js"
)

func prepareAttrs(vnode *Vnode, domNode js.Value) {
	if vnode.Attrs != nil && vnode.Attrs.Props != nil {
		for k, v := range *vnode.Attrs.Props {
			domNode.Call("setAttribute", k, v)
		}
	}
}

func prepareEvents(vnode *Vnode, domNode js.Value) {
	if vnode.Attrs != nil && vnode.Attrs.Events != nil {
		for k, f := range *vnode.Attrs.Events {
			callback := js.NewCallback(f)
			domNode.Call("addEventListener", k, callback)
		}
	}
}

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
