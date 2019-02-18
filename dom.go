package vn

import (
	"syscall/js"
)

type DomNode struct {
	binding js.Value
}

var instance *DomNode

func getDocument() DomNode {
	if instance == nil {
		docNode := js.Global().Get("document")
		instance = &DomNode{docNode}
	}

	return *instance
}

func (node DomNode) querySelector(element string) DomNode {
	bindingNode := getDocument().binding.Call("querySelector", element)

	return DomNode{bindingNode}
}

func (node DomNode) appendChild(child DomNode) {
	node.binding.Call("appendChild", child.binding)
}

func (node DomNode) remove() {
	node.binding.Call("remove")
}

func (node DomNode) replaceWith(next DomNode) {
	node.binding.Call("replaceWith", next.binding)
}

func (node DomNode) createTextNode(value string) DomNode {
	textNode := getDocument().binding.Call("createTextNode", value)

	return DomNode{textNode}
}
