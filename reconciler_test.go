package vn_test

import (
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	vn "github.com/mfrachet/go-vdom-wasm"
	vnd "github.com/mfrachet/go-vdom-wasm/dom"
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

func handleClick(arg []js.Value) {}
func TestReconciler_CreateInstance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	vnode := vn.H("li", &vn.Attrs{Props: &vn.Props{"class": "navbar"}, Events: &vn.Ev{"click": nil}}, "Hello world")

	mockParent := mock.NewMockDomNode(ctrl)
	mockChild := mock.NewMockDomNode(ctrl)
	mockParent.EXPECT().CreateElement("li").Return(mockChild).Times(1)
	mockChild.EXPECT().SetAttribute("class", "navbar").Times(1)
	mockChild.EXPECT().AddEventListener("click", nil).Times(1)

	domNode := vn.CreateInstance(mockParent, vnode)

	assert.Equal(t, domNode, mockChild)
}

func TestReconciler_CreateIfNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	domElement := vnd.DomElement{}
	vnode := vn.H("div", nil, vn.Children{})
	vnode.SetElement(domElement)

	mockParent := mock.NewMockDomNode(ctrl)

	domNode := vn.CreateIfNotExist(mockParent, vnode)

	assert.Equal(t, domNode, domElement)
}
