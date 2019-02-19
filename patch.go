package vn

import (
	"fmt"

	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

func initializeApp(rootNodeID string, initialNode Node) {
	rootNode := vnd.GetDocument().QuerySelector(rootNodeID)
	initialNode.createElement()
	domNode := *initialNode.getElement()

	rootNode.AppendChild(domNode)
}

func updateElement(parent vnd.DomNode, newNode Node, oldNode Node, index int) {
	if newNode == nil {
		oldElement := *oldNode.getElement()
		oldElement.Remove()
	} else {
		newNode.createElement()

		if oldNode == nil {
			// Adding a new child to the tree
			newNode.createElement()
			parent.AppendChild(*newNode.getElement())
		} else if !newNode.isSame(oldNode) {
			fmt.Println("lol", newNode, oldNode)
			// Replacing two different children
			newNode.createElement()

			oldElement := *oldNode.getElement()
			newElement := *newNode.getElement()
			oldElement.ReplaceWith(newElement)
		} else {
			// handling children
			newChildrenCount := newNode.childrenCount()
			oldChildrenCount := oldNode.childrenCount()

			max := vnh.Max(newChildrenCount, oldChildrenCount)

			for i := 0; i < max; i++ {
				var oldChild Node
				var newChild Node

				if oldNode.childrenCount() > i {
					oldChild = oldNode.getChildren()[i]
				}

				if newNode.childrenCount() > i {
					newChild = newNode.getChildren()[i]
				}

				updateElement(parent.ChildNodes(index), newChild, oldChild, i)
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
