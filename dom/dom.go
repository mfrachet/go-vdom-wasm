package vnd

import (
	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
	"syscall/js"
)

type DomNode interface {
	QuerySelector(string) DomNode
	AppendChild(DomNode)
	Remove()
	ReplaceWith(DomNode)
	CreateTextNode(string) DomNode
	CreateElement(string) DomNode
	SetAttribute(string, string)
	AddEventListener(string, func(js.Value, []js.Value) interface{})
	ChildNodes(int) DomNode
	GetBinding() js.Value
	SetBinding(js.Value)
}

type DomElement struct {
	binding js.Value
}

var instance *DomElement

func GetDocument() DomNode {
	if vnh.IsNil(instance) {
		docNode := js.Global().Get("document")
		instance = &DomElement{docNode}
	}

	return *instance
}

func (node DomElement) GetBinding() js.Value {
	return node.binding
}

func (node DomElement) SetBinding(nextBinding js.Value) {
	node.binding = nextBinding
}

func (node DomElement) QuerySelector(element string) DomNode {
	bindingNode := GetDocument().GetBinding().Call("querySelector", element)

	return DomElement{bindingNode}
}

func (node DomElement) AppendChild(child DomNode) {
	node.GetBinding().Call("appendChild", child.GetBinding())
}

func (node DomElement) Remove() {
	node.GetBinding().Call("remove")
}

func (node DomElement) ReplaceWith(next DomNode) {
	node.GetBinding().Call("replaceWith", next.GetBinding())
}

func (node DomElement) CreateTextNode(value string) DomNode {
	textNode := GetDocument().GetBinding().Call("createTextNode", value)

	return DomElement{textNode}
}

func (node DomElement) CreateElement(tag string) DomNode {
	element := GetDocument().GetBinding().Call("createElement", tag)

	return DomElement{element}
}

func (node DomElement) SetAttribute(attr string, value string) {
	node.GetBinding().Call("setAttribute", attr, value)
}

func (node DomElement) AddEventListener(eventName string, callback func(js.Value, []js.Value) interface{}) {
	node.GetBinding().Call("addEventListener", eventName, js.FuncOf(callback))
}

func (node DomElement) ChildNodes(index int) DomNode {
	element := node.GetBinding().Get("childNodes").Index(index)

	return DomElement{element}
}
