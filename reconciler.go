package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
)

func Append(domNode vnd.DomNode, virtualNode Node) {
	element := virtualNode.MakeDomNode(domNode)

	domNode.AppendChild(*element)
}
