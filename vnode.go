package vn

import (
	"fmt"
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

func (vnode *Vnode) hashCode() string {
	if vnode.Attrs != nil && vnode.Attrs.Props != nil {
		return fmt.Sprintf("%s/%v", vnode.TagName, *vnode.Attrs.Props)
	}

	return fmt.Sprintf("%s/%v", vnode.TagName, Attrs{})
}

func (vnode *Vnode) isSame(other Node) bool {
	return vnode.hashCode() == other.hashCode()
}

func (vnode *Vnode) childrenCount() int {
	switch vnode.Children.(type) {
	case Children:
		return len(vnode.Children.(Children))
	}

	return 0
}

func (vnode *Vnode) getChildren() interface{} {
	return vnode.Children
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

func (vnode *Vnode) getElement() *js.Value {
	return vnode.Element
}
