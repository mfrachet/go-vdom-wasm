package vn

import (
	"fmt"
)

type Children []*Vnode

type Vnode struct {
	TagName  string
	Attrs    *Attrs
	Children Children
	Text     *TextNode
	Element  *DomNode
}

func (vnode *Vnode) isSame(other Node) bool {
	if vnode.Text == nil {
		if other.getText() == nil {
			return vnode.hashCode() == other.hashCode()
		}

		return false
	}

	if other.getText() != nil {
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

func (vnode *Vnode) getElement() *DomNode {
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
		document := getDocument()
		domNode := document.createElement(vnode.TagName)

		if vnode.Attrs != nil {
			if vnode.Attrs.Props != nil {
				for attr, attrValue := range *vnode.Attrs.Props {
					domNode.setAttribute(attr, attrValue)
				}
			}

			if vnode.Attrs.Events != nil {
				for eventName, handler := range *vnode.Attrs.Events {
					domNode.addEventListener(eventName, handler)
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
		vnode.Element.appendChild(*textNode.getElement())
	} else {
		for _, el := range vnode.Children {
			el.createElement()
			vnode.Element.appendChild(*el.Element)
		}
	}
}
