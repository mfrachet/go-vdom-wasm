package vn_dom

import (
	"syscall/js"
)

type DomNode interface {
	GetDocument() DomNode
	QuerySelector(string) DomNode
	AppendChild(string) DomNode
	Remove()
	ReplaceWith(DomNode)
	CreateTextNode(string) DomNode
	CreateElement(string) DomNode
	SetAttribute(string, string)
	AddEventListener(string, func([]js.Value))
	ChildNodes(int)
}

type DomElement struct {
	binding js.Value
}

var instance *DomElement

func GetDocument() DomElement {
	if instance == nil {
		docNode := js.Global().Get("document")
		instance = &DomElement{docNode}
	}

	return *instance
}

func (node DomElement) QuerySelector(element string) DomElement {
	bindingNode := GetDocument().binding.Call("querySelector", element)

	return DomElement{bindingNode}
}

func (node DomElement) AppendChild(child DomElement) {
	node.binding.Call("appendChild", child.binding)
}

func (node DomElement) Remove() {
	node.binding.Call("remove")
}

func (node DomElement) ReplaceWith(next DomElement) {
	node.binding.Call("replaceWith", next.binding)
}

func (node DomElement) CreateTextNode(value string) DomElement {
	textNode := GetDocument().binding.Call("createTextNode", value)

	return DomElement{textNode}
}

func (node DomElement) CreateElement(tag string) DomElement {
	element := GetDocument().binding.Call("createElement", tag)

	return DomElement{element}
}

func (node DomElement) SetAttribute(attr string, value string) {
	node.binding.Call("setAttribute", attr, value)
}

func (node DomElement) AddEventListener(eventName string, callback func([]js.Value)) {
	node.binding.Call("addEventListener", eventName, js.NewCallback(callback))
}

func (node DomElement) ChildNodes(index int) DomElement {
	element := node.binding.Get("childNodes").Index(index)

	return DomElement{element}
}
