package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

func AppendToNode(domNode vnd.DomNode, virtualNode Node) {
	virtualNode.createElement()
	element := *virtualNode.getElement()

	domNode.AppendChild(element)
}

func updateElement(parent vnd.DomNode, newNode Node, oldNode Node, index int) {
	if vnh.IsNil(newNode) {
		oldElement := *oldNode.getElement()
		oldElement.Remove()
	} else {
		newNode.createElement()

		if vnh.IsNil(oldNode) {
			// Adding a new child to the tree
			AppendToNode(parent, newNode)
		} else if !newNode.isSame(oldNode) {
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
		rootNodeID := oldNodeRef.(string)
		rootNode := vnd.GetDocument().QuerySelector(rootNodeID)
		AppendToNode(rootNode, newVnode)
	default:
		oldVnode := oldNodeRef.(Node)

		updateElement(*oldVnode.getElement(), newVnode, oldVnode, 0)
	}
}
