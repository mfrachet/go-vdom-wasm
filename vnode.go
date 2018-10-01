package vn

import (
	"reflect"
	"syscall/js"
)

type Ev = map[string]func([]js.Value)
type Props = map[string]string
type Attrs struct {
	Props  *Props
	Events *Ev
}

type Vnode struct {
	TagName  string
	Attrs    *Attrs
	Children interface{}
	Element  *js.Value
}

func (vnode *Vnode) isSame(otherVnode *Vnode) bool {
	if vnode.Attrs == nil && otherVnode.Attrs == nil {
		return vnode.TagName == otherVnode.TagName
	}

	if vnode.Attrs.Props == nil && otherVnode.Attrs.Props == nil {
		return vnode.TagName == otherVnode.TagName
	}

	return vnode.TagName == otherVnode.TagName && reflect.DeepEqual(vnode.Attrs.Props, otherVnode.Attrs.Props)
}

func (vnode *Vnode) createElement() {
	if vnode.Element == nil {
		document := js.Global().Get("document")
		domNode := document.Call("createElement", vnode.TagName)

		if vnode.Attrs != nil {
			if vnode.Attrs.Props != nil {
				for attr, attrValue := range *vnode.Attrs.Props {
					domNode.Call("setAttribute", attr, attrValue)
				}
			}

			if vnode.Attrs.Events != nil {
				for eventName, handler := range *vnode.Attrs.Events {
					callback := js.NewCallback(handler)
					domNode.Call("addEventListener", eventName, callback)
				}
			}
		}

		vnode.Element = &domNode
		vnode.computeChildren()
	}
}

func (vnode *Vnode) computeChildren() {
	switch vnode.Children.(type) {
	case *TextNode:
		textNode := vnode.Children.(*TextNode)
		textNode.createElement()
		vnode.Element.Call("appendChild", *textNode.Element)
	case Children:
		for _, el := range vnode.Children.(Children) {
			el.createElement()
			vnode.Element.Call("appendChild", *el.Element)
		}
	}
}
