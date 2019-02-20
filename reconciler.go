package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

func Append(parent vnd.DomNode, child vnd.DomNode) {
	parent.AppendChild(child)
}

func Remove(domNode vnd.DomNode) {
	domNode.Remove()
}

func CreateText(document vnd.DomNode, virtualNode TextNode) vnd.DomElement {
	return document.CreateTextNode(virtualNode.Value)
}

func CreateInstance(document vnd.DomNode, vnode Node) vnd.DomElement {
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

		vnode.SetElement(domNode)
		computeChildren(document, vnode)
	}

	return *vnode.GetElement()
}

func computeChildren(document vnd.DomNode, vnode Node) {
	textNode := vnode.GetText()
	if vnh.NotNil(textNode) {
		textElement := CreateText(document, *textNode)
		textNode.SetElement(textElement)

		Append(vnode.GetElement(), textElement)
	} else {
		for _, el := range vnode.GetChildren() {
			el.MakeDomNode(document)
			vnode.GetElement().AppendChild(*el.GetElement())
		}
	}
}
