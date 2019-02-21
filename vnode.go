package vn

import (
	"fmt"

	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

type Children []*Vnode

type Vnode struct {
	TagName  string
	Attrs    *Attrs
	Children Children
	Text     *TextNode
	Element  *vnd.DomElement
	key      *string
}

func NewVNode(tagName string, attrs *Attrs, children Children, text *TextNode, element *vnd.DomElement, key *string) *Vnode {
	return &Vnode{tagName, attrs, children, text, element, key}
}

func (vnode *Vnode) IsSame(other Node) bool {
	currKey := vnode.GetKey()
	otherKey := other.GetKey()

	if currKey != nil && otherKey != nil {
		fmt.Println("They are the same => ", *currKey)
		return *currKey == *otherKey
	}

	if vnh.IsNil(vnode.Text) {
		if vnh.IsNil(other.GetText()) {
			return vnode.HashCode() == other.HashCode()
		}

		return false
	}

	if vnh.NotNil(other.GetText()) {
		hasSameText := other.GetText().IsSame(vnode.Text)
		return hasSameText && vnode.HashCode() == other.HashCode()
	}

	return false
}

func (vnode *Vnode) ChildrenCount() int {
	return len(vnode.Children)
}

func (vnode *Vnode) ChildAt(index int) Node {
	size := len(vnode.Children)

	if size > index {
		return vnode.Children[index]
	}

	return nil
}

func (vnode *Vnode) GetKey() *string {
	return vnode.key
}

func (vnode *Vnode) GetText() *TextNode {
	return vnode.Text
}

func (vnode *Vnode) GetElement() *vnd.DomElement {
	return vnode.Element
}

func (vnode *Vnode) HasElement() bool {
	return vnh.NotNil(vnode.GetElement())
}

func (vnode *Vnode) GetTagName() string {
	return vnode.TagName
}

func (vnode *Vnode) GetAttrs() *Attrs {
	return vnode.Attrs
}

func (vnode *Vnode) GetChildren() Children {
	return vnode.Children
}

func (vnode *Vnode) SetElement(element vnd.DomElement) {
	vnode.Element = &element
}

func (vnode *Vnode) HashCode() string {
	return fmt.Sprintf("%s/%v", vnode.TagName, *vnode.Attrs.Props)
}
