package vn

import "syscall/js"

type TextNode struct {
	Value   string
	Element *js.Value
}

func (textNode *TextNode) createElement() {
	if textNode.Element == nil {
		document := js.Global().Get("document")
		domNode := document.Call("createTextNode", textNode.Value)
		textNode.Element = &domNode
	}
}
