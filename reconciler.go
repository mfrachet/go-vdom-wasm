package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
)

func Append(domNode vnd.DomNode, virtualNode Node) {
	element := virtualNode.MakeDomNode(domNode)

	domNode.AppendChild(*element)
}

func Remove(domNode vnd.DomNode) {
	domNode.Remove()
}

func CreateInstance(document vnd.DomNode, virtualNode TextNode) vnd.DomElement {
	return document.CreateTextNode(virtualNode.Value)
}
