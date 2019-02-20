package vn_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	vn "github.com/mfrachet/go-vdom-wasm"
	"github.com/mfrachet/go-vdom-wasm/mock"
)

func TestReconciler_Append(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockParentNode := mock.NewMockDomNode(ctrl)
	mockChildNode := mock.NewMockDomNode(ctrl)

	mockParentNode.EXPECT().AppendChild(mockChildNode).Times(1)

	vn.Append(mockParentNode, mockChildNode)
}

func TestReconciler_Remove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDNode := mock.NewMockDomNode(ctrl)

	mockDNode.EXPECT().Remove().Times(1)

	vn.Remove(mockDNode)
}

func TestReconciler_CreateText(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	textNode := vn.TextNode{"Hello world", nil}
	mockDNode := mock.NewMockDomNode(ctrl)

	mockDNode.EXPECT().CreateTextNode("Hello world").Times(1)

	vn.CreateText(mockDNode, textNode)
}
