package vn

import (
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

func AppendToNode(domNode vnd.DomNode, virtualNode Node) {
	element := virtualNode.MakeDomNode(domNode)

	domNode.AppendChild(*element)
}

func updateElement(parent vnd.DomNode, newNode Node, oldNode Node, index int) {
	if vnh.IsNil(newNode) {
		oldElement := *oldNode.GetElement()
		oldElement.Remove()
	} else {
		if vnh.IsNil(oldNode) {
			// Adding a new child to the tree
			AppendToNode(parent, newNode)
		} else if !newNode.IsSame(oldNode) {
			// Replacing two different children
			newElement := *newNode.MakeDomNode(parent)

			oldElement := *oldNode.GetElement()
			oldElement.ReplaceWith(newElement)
		} else {
			// handling children
			newChildrenCount := newNode.ChildrenCount()
			oldChildrenCount := oldNode.ChildrenCount()

			max := vnh.Max(newChildrenCount, oldChildrenCount)

			for i := 0; i < max; i++ {

				oldChild := oldNode.ChildAt(i)
				newChild := newNode.ChildAt(i)

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

		updateElement(*oldVnode.GetElement(), newVnode, oldVnode, 0)
	}
}
