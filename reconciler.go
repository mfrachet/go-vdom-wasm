package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
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
