package vn_test

import (
	"syscall/js"
	"testing"

	"github.com/golang/mock/gomock"
	vn "github.com/mfrachet/go-vdom-wasm"
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	"github.com/mfrachet/go-vdom-wasm/mock"
)

func TestAppendChild(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVNode := mock.NewMockNode(ctrl)
	mockDNode := mock.NewMockDomNode(ctrl)

	el := vnd.DomElement{}
	el.SetBinding(js.ValueOf("Hello world"))

	mockVNode.EXPECT().CreateElement().Times(1)
	mockVNode.EXPECT().GetElement().Return(&el).Times(1)
	mockDNode.EXPECT().AppendChild(el).Times(1)

	vn.AppendToNode(mockDNode, mockVNode)
}
