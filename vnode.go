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

func NewVNode(tagName string, attrs *Attrs, children Children, text *TextNode, element *vnd.DomElement) *Vnode {
	return &Vnode{tagName, attrs, children, text, element}
}

func (vnode *Vnode) IsSame(other Node) bool {
	if vnh.IsNil(vnode.Text) {
		if vnh.IsNil(other.GetText()) {
			return vnode.HashCode() == other.HashCode()
		}

		return false
	}

	if vnh.NotNil(other.GetText()) {
		return vnode.Text.Value == other.GetText().Value && vnode.HashCode() == other.HashCode()
	}

	return false
}

func (vnode *Vnode) ChildrenCount() int {
	return len(vnode.Children)
}

func (vnode *Vnode) GetChildren() Children {
	return vnode.Children
}

func (vnode *Vnode) GetText() *TextNode {
	return vnode.Text
}

func (vnode *Vnode) GetElement() *vnd.DomElement {
	return vnode.Element
}

func (vnode *Vnode) HashCode() string {
	if vnh.NotNil(vnode.Attrs) && vnh.NotNil(vnode.Attrs.Props) {
		return fmt.Sprintf("%s/%v", vnode.TagName, *vnode.Attrs.Props)
	}

	return fmt.Sprintf("%s/%v", vnode.TagName, Attrs{})
}

func (vnode *Vnode) CreateElement() {
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
		textNode.CreateElement()
		vnode.Element.AppendChild(*textNode.GetElement())
	} else {
		for _, el := range vnode.Children {
			el.CreateElement()
			vnode.Element.AppendChild(*el.Element)
		}
	}
}
