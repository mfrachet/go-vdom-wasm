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

func CreateText(parent vnd.DomNode, virtualNode TextNode) vnd.DomElement {
	return parent.CreateTextNode(virtualNode.Value)
}

func CreateInstance(parent vnd.DomNode, vnode Node) vnd.DomElement {
	if vnh.IsNil(vnode.GetElement()) {
		domNode := parent.CreateElement(vnode.GetTagName())
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
		ComputeChildren(parent, vnode)
	}

	return *vnode.GetElement()
}

func ComputeChildren(parent vnd.DomNode, vnode Node) {
	textNode := vnode.GetText()
	if vnh.NotNil(textNode) {
		textElement := CreateText(parent, *textNode)
		textNode.SetElement(textElement)

		Append(vnode.GetElement(), textElement)
	} else {
		for _, el := range vnode.GetChildren() {
			CreateInstance(parent, el)
			vnode.GetElement().AppendChild(*el.GetElement())
		}
	}
}
