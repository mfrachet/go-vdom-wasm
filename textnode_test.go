package vn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTextNOde(t *testing.T) {
	a := &TextNode{Value: "Hello world"}
	b := &TextNode{Value: "Hello world"}
	c := &TextNode{Value: "Hello world2"}

	assert.Equal(t, true, a.isSame(b))
	assert.Equal(t, false, a.isSame(c))
}
