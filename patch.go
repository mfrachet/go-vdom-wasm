package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

func updateElement(parent vnd.DomNode, newNode Node, oldNode Node) {
	if vnh.IsNil(newNode) {
		Remove(*oldNode.GetElement())
	} else if vnh.IsNil(oldNode) {
		newElement := createIfNotExist(parent, newNode)

		Append(parent, newElement)
	} else if !newNode.IsSame(oldNode) {
		newElement := createIfNotExist(parent, newNode)

		oldElement := *oldNode.GetElement()
		oldElement.ReplaceWith(newElement)
	} else {
		newNode.SetElement(*oldNode.GetElement())

		newChildrenCount := newNode.ChildrenCount()
		oldChildrenCount := oldNode.ChildrenCount()

		max := vnh.Max(newChildrenCount, oldChildrenCount)

		for i := 0; i < max; i++ {
			oldChild := oldNode.ChildAt(i)
			newChild := newNode.ChildAt(i)

			updateElement(*oldNode.GetElement(), newChild, oldChild)
		}
	}
}

func Patch(oldNodeRef interface{}, newVnode Node) {
	switch oldNodeRef.(type) {
	case string:
		rootNodeID := oldNodeRef.(string)
		rootNode := vnd.GetDocument().QuerySelector(rootNodeID)

		newElement := createIfNotExist(rootNode, newVnode)

		Append(rootNode, newElement)
	default:
		oldVnode := oldNodeRef.(Node)

		updateElement(*oldVnode.GetElement(), newVnode, oldVnode)
	}
}
