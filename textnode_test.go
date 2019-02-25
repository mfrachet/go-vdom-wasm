package vn_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	vn "github.com/mfrachet/go-vdom-wasm"
	"github.com/mfrachet/go-vdom-wasm/mock"
	"github.com/stretchr/testify/assert"
)

func TestTextNode_IsSame(t *testing.T) {
	a := vn.NewTextnode("Hello world")
	b := vn.NewTextnode("Hello world")
	c := vn.NewTextnode("Hello world2")

	assert.Equal(t, true, a.IsSame(b))
	assert.Equal(t, false, a.IsSame(c))
}

func TestTextNode_Accessors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	a := vn.NewTextnode("Hello world")

	mockDomNode := mock.NewMockDomNode(ctrl)
	a.SetElement(mockDomNode)

	assert.Equal(t, mockDomNode, *a.GetElement())
	assert.Equal(t, "Hello world", a.GetValue())
}
