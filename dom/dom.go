package vnd

import (
	"syscall/js"

	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

type DomNode interface {
	QuerySelector(string) DomElement
	AppendChild(DomNode)
	Remove()
	ReplaceWith(DomElement)
	CreateTextNode(string) DomElement
	CreateElement(string) DomElement
	SetAttribute(string, string)
	AddEventListener(string, func([]js.Value))
	ChildNodes(int) DomElement
	GetBinding() js.Value
	SetBinding(js.Value)
	GetParent() DomElement
}

type DomElement struct {
	binding js.Value
}

var instance *DomElement

func GetDocument() DomElement {
	if vnh.IsNil(instance) {
		docNode := js.Global().Get("document")
		instance = &DomElement{docNode}
	}

	return *instance
}

func (node DomElement) GetParent() DomElement {
	parent := node.GetBinding().Get("parentElement")

	return DomElement{parent}
}

func (node DomElement) GetBinding() js.Value {
	return node.binding
}

func (node DomElement) SetBinding(nextBinding js.Value) {
	node.binding = nextBinding
}

func (node DomElement) QuerySelector(element string) DomElement {
	bindingNode := GetDocument().GetBinding().Call("querySelector", element)

	return DomElement{bindingNode}
}

func (node DomElement) AppendChild(child DomNode) {
	node.GetBinding().Call("appendChild", child.GetBinding())
}

func (node DomElement) Remove() {
	node.GetBinding().Call("remove")
}

func (node DomElement) ReplaceWith(next DomElement) {
	node.GetBinding().Call("replaceWith", next.GetBinding())
}

func (node DomElement) CreateTextNode(value string) DomElement {
	textNode := GetDocument().GetBinding().Call("createTextNode", value)

	return DomElement{textNode}
}

func (node DomElement) CreateElement(tag string) DomElement {
	element := GetDocument().GetBinding().Call("createElement", tag)

	return DomElement{element}
}

func (node DomElement) SetAttribute(attr string, value string) {
	node.GetBinding().Call("setAttribute", attr, value)
}

func (node DomElement) AddEventListener(eventName string, callback func([]js.Value)) {
	node.GetBinding().Call("addEventListener", eventName, js.NewCallback(callback))
}

func (node DomElement) ChildNodes(index int) DomElement {
	element := node.GetBinding().Get("childNodes").Index(index)

	return DomElement{element}
}
