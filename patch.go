package vn

import (
	"syscall/js"
)

func initializeApp(rootNodeID string, initialNode *Vnode) {
	rootNode := js.Global().Get("document").Call("querySelector", rootNodeID)
	initialNode.createElement()
	domNode := *initialNode.Element

	rootNode.Call("appendChild", domNode)
}

func updateElement(parent js.Value, newNode *Vnode, oldNode *Vnode, index int) {
	if oldNode == nil {
		newNode.createElement()
		parent.Call("appendChild", newNode.Element)
	} else if newNode == nil {
		indexToRemove := parent.Get("childNodes").Index(index).Int()
		parent.Call("removeChild", indexToRemove)
	} else if !newNode.isSame(oldNode) {
		newNode.createElement()
		oldNode.Element.Call("replaceWith", *newNode.Element)
	} else {
		newNode.createElement()
		switch oldNode.Children.(type) {
		case Children:
			oldChildren := oldNode.Children.(Children)
			newChildren := newNode.Children.(Children)
			oldChildLength := len(oldChildren)
			newChildLength := len(newChildren)

			for i := 0; i < oldChildLength || i < newChildLength; i++ {
				updateElement(parent.Get("childNodes").Index(i), newChildren[i], oldChildren[i], i)
			}
		case *TextNode:
			oldTextNode := oldNode.Children.(*TextNode)
			newTextNode := newNode.Children.(*TextNode)

			if oldTextNode.Value != newTextNode.Value {
				newTextNode.createElement()
				oldNode.Element.Call("replaceWith", *newNode.Element)
			}

		}
	}
}

func Patch(oldNodeRef interface{}, newVnode *Vnode) {
	switch oldNodeRef.(type) {
	case string:
		initializeApp(oldNodeRef.(string), newVnode)
	default:
		oldVnode := oldNodeRef.(*Vnode)

		updateElement(*oldVnode.Element, newVnode, oldVnode, 0)
	}
}
