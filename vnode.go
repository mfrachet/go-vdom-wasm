package vn

import (
	"fmt"

	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

type Children []Node

type Node interface {
	GetElement() *vnd.DomNode
	GetTagName() string
	GetAttrs() *Attrs
	GetKey() *string
	HasElement() bool
	SetElement(vnd.DomNode)
	HashCode() string
	ChildrenCount() int
	ChildAt(int) Node
	GetText() *TextNode
	GetChildren() Children
	IsSame(Node) bool
}

type Vnode struct {
	tagname  string
	attrs    *Attrs
	children Children
	text     *TextNode
	element  *vnd.DomNode
	key      KeyIdentifier
}

func NewNode(tagname string, attrs *Attrs, children Children, text *TextNode, element *vnd.DomNode, key KeyIdentifier) Node {
	return &Vnode{tagname, attrs, children, text, element, key}
}

func (vnode *Vnode) IsSame(other Node) bool {
	currKey := vnode.GetKey()
	otherKey := other.GetKey()

	if currKey != nil && otherKey != nil {
		return *currKey == *otherKey
	}

	if vnh.IsNil(vnode.text) {
		if vnh.IsNil(other.GetText()) {
			return vnode.HashCode() == other.HashCode()
		}

		return false
	}

	if vnh.NotNil(other.GetText()) {
		hasSameText := other.GetText().IsSame(vnode.text)
		return hasSameText && vnode.HashCode() == other.HashCode()
	}

	return false
}

func (vnode *Vnode) ChildrenCount() int {
	return len(vnode.children)
}

func (vnode *Vnode) ChildAt(index int) Node {
	size := len(vnode.children)

	if size > index {
		return vnode.children[index]
	}

	return nil
}

func (vnode *Vnode) GetKey() *string {
	if vnh.IsNil(vnode.key) {
		return nil
	}

	key := vnode.key.GetValue()
	return &key
}

func (vnode *Vnode) GetText() *TextNode {
	return vnode.text
}

func (vnode *Vnode) GetElement() *vnd.DomNode {
	return vnode.element
}

func (vnode *Vnode) HasElement() bool {
	return vnh.NotNil(vnode.GetElement())
}

func (vnode *Vnode) GetTagName() string {
	return vnode.tagname
}

func (vnode *Vnode) GetAttrs() *Attrs {
	return vnode.attrs
}

func (vnode *Vnode) GetChildren() Children {
	return vnode.children
}

func (vnode *Vnode) SetElement(element vnd.DomNode) {
	vnode.element = &element
}

func (vnode *Vnode) HashCode() string {
	return fmt.Sprintf("%s/%v", vnode.tagname, *vnode.attrs.Props)
}
