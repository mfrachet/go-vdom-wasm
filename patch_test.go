package vn_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	vn "github.com/mfrachet/go-vdom-wasm"
	"github.com/mfrachet/go-vdom-wasm/mock"
)

func TestAppendChild(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVNode := mock.NewMockNode(ctrl)
	mockDNode := mock.NewMockDomNode(ctrl)

	mockVNode.EXPECT().CreateElement().Times(1)

	vn.AppendToNode(mockDNode, mockVNode)

}
