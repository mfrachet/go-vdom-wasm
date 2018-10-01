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
	document := js.Global().Get("document")

	switch node.(type) {
	case *TextNode:
		vnode := node.(*TextNode)
		domNode := document.Call("createTextNode", vnode.Value)
		vnode.Element = &domNode

		return vnode.Element

	default:
		vnode := node.(*Vnode)
		domNode := document.Call("createElement", vnode.TagName)

		prepareAttrs(vnode, domNode)
		prepareEvents(vnode, domNode)

		switch vnode.Children.(type) {
		case *TextNode:
			domNode.Call("appendChild", *CreateElement(vnode.Children))
		case Children:
			for _, el := range vnode.Children.(Children) {
				childNode := *CreateElement(el)
				domNode.Call("appendChild", childNode)
			}
		}

		vnode.Element = &domNode

		return vnode.Element
	}
}
