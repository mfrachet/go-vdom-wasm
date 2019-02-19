package vn_test

import (
	"syscall/js"
	"testing"

	"github.com/golang/mock/gomock"
	vn "github.com/mfrachet/go-vdom-wasm"
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	"github.com/mfrachet/go-vdom-wasm/mock"
)

func TestReconciler_Append(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVNode := mock.NewMockNode(ctrl)
	mockDNode := mock.NewMockDomNode(ctrl)

	el := vnd.DomElement{}
	el.SetBinding(js.ValueOf("To append"))

	mockVNode.EXPECT().MakeDomNode(mockDNode).Return(&el).Times(1)
	mockDNode.EXPECT().AppendChild(el).Times(1)

	vn.Append(mockDNode, mockVNode)
}
