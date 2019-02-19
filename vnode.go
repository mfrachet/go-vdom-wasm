package vn

import (
	"fmt"

	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

type Children []*Vnode

type Vnode struct {
	TagName  string
	Attrs    *Attrs
	Children Children
	Text     *TextNode
	Element  *vnd.DomElement
}

func (vnode *Vnode) isSame(other Node) bool {
	if vnh.IsNil(vnode.Text) {
		if vnh.IsNil(other.getText()) {
			return vnode.hashCode() == other.hashCode()
		}

		return false
	}

	if vnh.NotNil(other.getText()) {
		return vnode.Text.Value == other.getText().Value && vnode.hashCode() == other.hashCode()
	}

	return false
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

func (vnode *Vnode) getElement() *vnd.DomElement {
	return vnode.Element
}

func (vnode *Vnode) hashCode() string {
	if vnh.NotNil(vnode.Attrs) && vnh.NotNil(vnode.Attrs.Props) {
		return fmt.Sprintf("%s/%v", vnode.TagName, *vnode.Attrs.Props)
	}

	return fmt.Sprintf("%s/%v", vnode.TagName, Attrs{})
}

func (vnode *Vnode) createElement() {
	if vnh.IsNil(vnode.Element) {
		document := vnd.GetDocument()
		domNode := document.CreateElement(vnode.TagName)

		if vnh.NotNil(vnode.Attrs) {
			if vnh.NotNil(vnode.Attrs.Props) {
				for attr, attrValue := range *vnode.Attrs.Props {
					domNode.SetAttribute(attr, attrValue)
				}
			}

			if vnh.NotNil(vnode.Attrs.Events) {
				for eventName, handler := range *vnode.Attrs.Events {
					domNode.AddEventListener(eventName, handler)
				}
			}
		}

		vnode.Element = &domNode
		vnode.computeChildren()
	}
}

func (vnode *Vnode) computeChildren() {
	if vnh.NotNil(vnode.Text) {
		textNode := vnode.Text
		textNode.createElement()
		vnode.Element.AppendChild(*textNode.getElement())
	} else {
		for _, el := range vnode.Children {
			el.createElement()
			vnode.Element.AppendChild(*el.Element)
		}
	}
}
