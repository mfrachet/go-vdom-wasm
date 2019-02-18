package vn

type Node interface {
	createElement()
	getElement() *DomNode
	hashCode() string
	childrenCount() int
	getChildren() Children
	getText() *TextNode
	isSame(Node) bool
}
