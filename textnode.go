package vn

type TextNode struct {
	Value   string
	Element *DomNode
}

func (textNode *TextNode) createElement() {
	if textNode.Element == nil {
		document := getDocument()
		domNode := document.createTextNode(textNode.Value)
		textNode.Element = &domNode
	}
}

func (textNode *TextNode) getElement() *DomNode {
	return textNode.Element
}

func (textNode *TextNode) hashCode() string {
	return textNode.Value
}

func (textNode *TextNode) isSame(other Node) bool {
	return textNode.hashCode() == other.hashCode()
}

func (textNode *TextNode) childrenCount() int {
	return -1
}

func (textNode *TextNode) getChildren() Children {
	return nil
}

func (textNode *TextNode) getText() *TextNode {
	return textNode
}
