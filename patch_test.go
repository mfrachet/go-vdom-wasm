package vn_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	vn "github.com/mfrachet/go-vdom-wasm"
	"github.com/mfrachet/go-vdom-wasm/mock"
)

func TestAppendChild(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockVNode := mock.NewMockNode(ctrl)
	mockDNode := mock.NewMockDomNode(ctrl)

	vn.AppendToNode(mockDNode, mockVNode)
}
