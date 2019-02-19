package vnh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNil(t *testing.T) {
	var node *[]int

	str := "Hello world"
	pointer := &str

	assert.Equal(t, true, IsNil(nil))
	assert.Equal(t, true, IsNil(node))
	assert.Equal(t, false, IsNil(pointer))
}

func TestNotNil(t *testing.T) {
	var node *[]int

	str := "Hello world"
	pointer := &str

	assert.Equal(t, false, NotNil(nil))
	assert.Equal(t, true, NotNil(pointer))
	assert.Equal(t, false, NotNil(node))
}
