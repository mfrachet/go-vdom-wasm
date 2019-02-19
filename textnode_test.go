package vn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextNode_IsSame(t *testing.T) {
	a := &TextNode{Value: "Hello world"}
	b := &TextNode{Value: "Hello world"}
	c := &TextNode{Value: "Hello world2"}

	assert.Equal(t, true, a.IsSame(b))
	assert.Equal(t, false, a.IsSame(c))
}

func TestTextNode_ChildrenCount(t *testing.T) {
	a := &TextNode{Value: "Hello world"}
	assert.Equal(t, -1, a.ChildrenCount())
}

func TestTextNode_ChildAt(t *testing.T) {
	a := TextNode{Value: "Hello world"}
	child := a.ChildAt(1)

	assert.Equal(t, nil, child)
}
