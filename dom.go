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
