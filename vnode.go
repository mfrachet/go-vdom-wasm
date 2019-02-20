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

func (vnode *Vnode) ChildAt(index int) Node {
	size := len(vnode.Children)

	if size > index {
		return vnode.Children[index]
	}

	return nil
}

func (vnode *Vnode) GetText() *TextNode {
	return vnode.Text
}

func (vnode *Vnode) GetElement() *vnd.DomElement {
	return vnode.Element
}

func (vnode *Vnode) GetTagName() string {
	return vnode.TagName
}

func (vnode *Vnode) GetAttrs() *Attrs {
	return vnode.Attrs
}

func (vnode *Vnode) GetChildren() Children {
	return vnode.Children
}

func (vnode *Vnode) SetElement(element vnd.DomElement) {
	vnode.Element = &element
}

func (vnode *Vnode) HashCode() string {
	if vnh.NotNil(vnode.Attrs) && vnh.NotNil(vnode.Attrs.Props) {
		return fmt.Sprintf("%s/%v", vnode.TagName, *vnode.Attrs.Props)
	}

	return fmt.Sprintf("%s/%v", vnode.TagName, Attrs{})
}

func (vnode *Vnode) MakeDomNode(document vnd.DomNode) *vnd.DomElement {
	if vnh.IsNil(vnode.GetElement()) {
		domNode := document.CreateElement(vnode.GetTagName())
		attrs := vnode.GetAttrs()

		if vnh.NotNil(attrs) {
			if vnh.NotNil(attrs.Props) {
				for attr, attrValue := range *attrs.Props {
					domNode.SetAttribute(attr, attrValue)
				}
			}

			if vnh.NotNil(attrs.Events) {
				for eventName, handler := range *attrs.Events {
					domNode.AddEventListener(eventName, handler)
				}
			}
		}

		vnode.Element = &domNode
		vnode.computeChildren(document)
	}

	return vnode.Element
}

func (vnode *Vnode) computeChildren(document vnd.DomNode) {
	textNode := vnode.Text
	if vnh.NotNil(textNode) {
		textElement := CreateText(document, *textNode)
		textNode.SetElement(textElement)

		Append(vnode.GetElement(), textElement)
	} else {
		for _, el := range vnode.Children {
			el.MakeDomNode(document)
			vnode.Element.AppendChild(*el.Element)
		}
	}
}
