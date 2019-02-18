package vn

func initializeApp(rootNodeID string, initialNode Node) {
	rootNode := getDocument().querySelector(rootNodeID)
	initialNode.createElement()
	domNode := *initialNode.getElement()

	rootNode.appendChild(domNode)
}

func updateElement(parent DomNode, newNode Node, oldNode Node, index int) {

	if oldNode == nil {
		// Adding a new child to the tree
		newNode.createElement()
		parent.appendChild(*newNode.getElement())
	} else if newNode == nil {
		oldNode.getElement().remove()
	} else if !newNode.isSame(oldNode) {
		// Replacing two different children
		newNode.createElement()
		oldNode.getElement().replaceWith(*newNode.getElement())
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

			oldChildren := oldNode.getChildren()
			newChildren := newNode.getChildren()

			for i := 0; i < max; i++ {
				var newChild Node
				var oldChild Node

				if oldChildrenCount > i {
					oldChild = oldChildren[i]
				}

				if newChildrenCount > i {
					newChild = newChildren[i]
				}

				if newChild != nil && newChild.childrenCount() > 0 {
					updateElement(parent.childNodes(index), newChild, oldChild, i)
				} else {
					updateElement(parent, newChild, oldChild, i)
				}

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
