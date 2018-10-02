package vn

import (
	"syscall/js"
)

func initializeApp(rootNodeID string, initialNode Node) {
	rootNode := js.Global().Get("document").Call("querySelector", rootNodeID)
	initialNode.createElement()
	domNode := *initialNode.getElement()

	rootNode.Call("appendChild", domNode)
}

func updateElement(parent js.Value, newNode Node, oldNode Node, index int) {

	if oldNode == nil {
		// Adding a new child to the tree
		newNode.createElement()
		parent.Call("appendChild", *newNode.getElement())
	} else if newNode == nil {
		oldNode.getElement().Call("remove")
	} else if !newNode.isSame(oldNode) {
		// Replacing two different children
		newNode.createElement()
		oldNode.getElement().Call("replaceWith", *newNode.getElement())
	} else {
		// handling children
		newChildrenCount := newNode.childrenCount()

		if newChildrenCount > 0 {
			newNode.createElement()
			oldChildrenCount := oldNode.childrenCount()

			var max int = newChildrenCount

			if oldChildrenCount > newChildrenCount {
				max = oldChildrenCount
			}

			oldChildren := oldNode.getChildren().(Children)
			newChildren := newNode.getChildren().(Children)

			for i := 0; i < max; i++ {
				var newChild Node
				var oldChild Node

				if oldChildrenCount > i {
					oldChild = oldChildren[i]
				}

				if newChildrenCount > i {
					newChild = newChildren[i]
				}

				updateElement(parent.Get("childNodes").Index(index+1), newChild, oldChild, i)
			}
		}
	}
}

func Patch(oldNodeRef interface{}, newVnode Node) {
	switch oldNodeRef.(type) {
	case string:
		initializeApp(oldNodeRef.(string), newVnode)
	default:
		oldVnode := oldNodeRef.(Node)

		updateElement(*oldVnode.getElement(), newVnode, oldVnode, 0)
	}
}
