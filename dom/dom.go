package vn_dom

import (
	"syscall/js"
)

type DomNode struct {
	binding js.Value
}

var instance *DomNode

func GetDocument() DomNode {
	if instance == nil {
		docNode := js.Global().Get("document")
		instance = &DomNode{docNode}
	}

	return *instance
}

func (node DomNode) QuerySelector(element string) DomNode {
	bindingNode := GetDocument().binding.Call("querySelector", element)

	return DomNode{bindingNode}
}

func (node DomNode) AppendChild(child DomNode) {
	node.binding.Call("appendChild", child.binding)
}

func (node DomNode) Remove() {
	node.binding.Call("remove")
}

func (node DomNode) ReplaceWith(next DomNode) {
	node.binding.Call("replaceWith", next.binding)
}

func (node DomNode) CreateTextNode(value string) DomNode {
	textNode := GetDocument().binding.Call("createTextNode", value)

	return DomNode{textNode}
}

func (node DomNode) CreateElement(tag string) DomNode {
	element := GetDocument().binding.Call("createElement", tag)

	return DomNode{element}
}

func (node DomNode) SetAttribute(attr string, value string) {
	node.binding.Call("setAttribute", attr, value)
}

func (node DomNode) AddEventListener(eventName string, callback func([]js.Value)) {
	node.binding.Call("addEventListener", eventName, js.NewCallback(callback))
}

func (node DomNode) ChildNodes(index int) DomNode {
	element := node.binding.Get("childNodes").Index(index)

	return DomNode{element}
}
