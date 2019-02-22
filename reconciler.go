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

func CreateText(parent vnd.DomNode, virtualNode TextNode) vnd.DomNode {
	return parent.CreateTextNode(virtualNode.Value)
}

func CreateInstance(parent vnd.DomNode, vnode Node) vnd.DomNode {
	domNode := parent.CreateElement(vnode.GetTagName())
	attrs := vnode.GetAttrs()

	for attr, attrValue := range *attrs.Props {
		domNode.SetAttribute(attr, attrValue)
	}

	for eventName, handler := range *attrs.Events {
		domNode.AddEventListener(eventName, handler)
	}

	return domNode
}

func ComputeChildren(parent vnd.DomNode, vnode Node) {
	textNode := vnode.GetText()
	if vnh.NotNil(textNode) {
		textElement := CreateText(parent, *textNode)
		textNode.SetElement(textElement)

		Append(parent, textElement)
	} else {
		for _, el := range vnode.GetChildren() {
			childElement := CreateIfNotExist(parent, el)
			parent.AppendChild(childElement)
		}
	}
}

func CreateIfNotExist(parent vnd.DomNode, vnode Node) vnd.DomNode {
	if vnode.HasElement() {
		return *vnode.GetElement()
	}

	newElement := CreateInstance(parent, vnode)
	vnode.SetElement(newElement)

	ComputeChildren(newElement, vnode)

	return newElement
}
