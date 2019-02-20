// Code generated by MockGen. DO NOT EDIT.
// Source: node.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	go_vdom_wasm "github.com/mfrachet/go-vdom-wasm"
	dom "github.com/mfrachet/go-vdom-wasm/dom"
	reflect "reflect"
)

// MockNode is a mock of Node interface
type MockNode struct {
	ctrl     *gomock.Controller
	recorder *MockNodeMockRecorder
}

// MockNodeMockRecorder is the mock recorder for MockNode
type MockNodeMockRecorder struct {
	mock *MockNode
}

// NewMockNode creates a new mock instance
func NewMockNode(ctrl *gomock.Controller) *MockNode {
	mock := &MockNode{ctrl: ctrl}
	mock.recorder = &MockNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNode) EXPECT() *MockNodeMockRecorder {
	return m.recorder
}

// MakeDomNode mocks base method
func (m *MockNode) MakeDomNode(arg0 dom.DomNode) *dom.DomElement {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeDomNode", arg0)
	ret0, _ := ret[0].(*dom.DomElement)
	return ret0
}

// MakeDomNode indicates an expected call of MakeDomNode
func (mr *MockNodeMockRecorder) MakeDomNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeDomNode", reflect.TypeOf((*MockNode)(nil).MakeDomNode), arg0)
}

// GetElement mocks base method
func (m *MockNode) GetElement() *dom.DomElement {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetElement")
	ret0, _ := ret[0].(*dom.DomElement)
	return ret0
}

// GetElement indicates an expected call of GetElement
func (mr *MockNodeMockRecorder) GetElement() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetElement", reflect.TypeOf((*MockNode)(nil).GetElement))
}

// GetTagName mocks base method
func (m *MockNode) GetTagName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTagName indicates an expected call of GetTagName
func (mr *MockNodeMockRecorder) GetTagName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagName", reflect.TypeOf((*MockNode)(nil).GetTagName))
}

// GetAttrs mocks base method
func (m *MockNode) GetAttrs() *go_vdom_wasm.Attrs {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttrs")
	ret0, _ := ret[0].(*go_vdom_wasm.Attrs)
	return ret0
}

// GetAttrs indicates an expected call of GetAttrs
func (mr *MockNodeMockRecorder) GetAttrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttrs", reflect.TypeOf((*MockNode)(nil).GetAttrs))
}

// SetElement mocks base method
func (m *MockNode) SetElement(arg0 dom.DomElement) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetElement", arg0)
}

// SetElement indicates an expected call of SetElement
func (mr *MockNodeMockRecorder) SetElement(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetElement", reflect.TypeOf((*MockNode)(nil).SetElement), arg0)
}

// HashCode mocks base method
func (m *MockNode) HashCode() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashCode")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashCode indicates an expected call of HashCode
func (mr *MockNodeMockRecorder) HashCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashCode", reflect.TypeOf((*MockNode)(nil).HashCode))
}

// ChildrenCount mocks base method
func (m *MockNode) ChildrenCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChildrenCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// ChildrenCount indicates an expected call of ChildrenCount
func (mr *MockNodeMockRecorder) ChildrenCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChildrenCount", reflect.TypeOf((*MockNode)(nil).ChildrenCount))
}

// ChildAt mocks base method
func (m *MockNode) ChildAt(arg0 int) go_vdom_wasm.Node {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChildAt", arg0)
	ret0, _ := ret[0].(go_vdom_wasm.Node)
	return ret0
}

// ChildAt indicates an expected call of ChildAt
func (mr *MockNodeMockRecorder) ChildAt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChildAt", reflect.TypeOf((*MockNode)(nil).ChildAt), arg0)
}

// GetText mocks base method
func (m *MockNode) GetText() *go_vdom_wasm.TextNode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetText")
	ret0, _ := ret[0].(*go_vdom_wasm.TextNode)
	return ret0
}

// GetText indicates an expected call of GetText
func (mr *MockNodeMockRecorder) GetText() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetText", reflect.TypeOf((*MockNode)(nil).GetText))
}

// GetChildren mocks base method
func (m *MockNode) GetChildren() go_vdom_wasm.Children {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChildren")
	ret0, _ := ret[0].(go_vdom_wasm.Children)
	return ret0
}

// GetChildren indicates an expected call of GetChildren
func (mr *MockNodeMockRecorder) GetChildren() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChildren", reflect.TypeOf((*MockNode)(nil).GetChildren))
}

// IsSame mocks base method
func (m *MockNode) IsSame(arg0 go_vdom_wasm.Node) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSame", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSame indicates an expected call of IsSame
func (mr *MockNodeMockRecorder) IsSame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSame", reflect.TypeOf((*MockNode)(nil).IsSame), arg0)
}
