package vn

import (
	"fmt"
	"syscall/js"
)

type Children []*Vnode

type Vnode struct {
	TagName  string
	Attrs    *Attrs
	Children Children
	Text     *TextNode
	Element  *js.Value
}

func (vnode *Vnode) isSame(other Node) bool {
	return vnode.Text == other.getText() && vnode.hashCode() == other.hashCode()
}

func (vnode *Vnode) childrenCount() int {
	return len(vnode.Children)
}

func (vnode *Vnode) getChildren() Children {
	return vnode.Children
}

func (vnode *Vnode) getText() *TextNode {
	return vnode.Text
}

func (vnode *Vnode) getElement() *js.Value {
	return vnode.Element
}

func (vnode *Vnode) hashCode() string {
	if vnode.Attrs != nil && vnode.Attrs.Props != nil {
		return fmt.Sprintf("%s/%v", vnode.TagName, *vnode.Attrs.Props)
	}

	return fmt.Sprintf("%s/%v", vnode.TagName, Attrs{})
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
	if vnode.Text != nil {
		textNode := vnode.Text
		textNode.createElement()
		vnode.Element.Call("appendChild", *textNode.getElement())
	} else {
		for _, el := range vnode.Children {
			el.createElement()
			vnode.Element.Call("appendChild", *el.Element)
		}
	}
}
