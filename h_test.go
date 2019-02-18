package vn_test

import (
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

import vn "github.com/mfrachet/go-vdom-wasm"

func TestHWithStringChildren(t *testing.T) {
	cases := []struct {
		TagName        string
		childrenString string
		want           *vn.Vnode
	}{
		{"div", "Hello world", &vn.Vnode{TagName: "div", Attrs: nil, Text: &vn.TextNode{Value: "Hello world", Element: nil}, Element: nil}},
		{"span", "Hello", &vn.Vnode{TagName: "span", Attrs: nil, Text: &vn.TextNode{Value: "Hello", Element: nil}, Element: nil}},
	}
	for _, c := range cases {
		got := vn.H(c.TagName, nil, c.childrenString)

		assert.Equal(t, c.want, got)
	}
}

func TestHWithChildren(t *testing.T) {
	childrenString := "Hello Children"

	expectedChild := &vn.Vnode{TagName: "div", Attrs: nil, Text: &vn.TextNode{childrenString, nil}, Element: nil}
	expectedVnode := &vn.Vnode{TagName: "div", Attrs: nil, Children: vn.Children{expectedChild}, Element: nil}

	currentVnode := vn.H("div", nil, vn.Children{vn.H("div", nil, childrenString)})

	assert.Equal(t, expectedVnode, currentVnode)
}

func TestHWithAttributes(t *testing.T) {
	ev := &vn.Ev{"f": func(args []js.Value) {}}
	cases := []struct {
		Attrs *vn.Attrs
		want  *vn.Vnode
	}{
		{&vn.Attrs{}, &vn.Vnode{TagName: "div", Attrs: &vn.Attrs{}, Text: &vn.TextNode{Value: "Hello world", Element: nil}, Element: nil}},
		{&vn.Attrs{Props: &vn.Props{"class": "hello"}}, &vn.Vnode{TagName: "div", Attrs: &vn.Attrs{Props: &vn.Props{"class": "hello"}}, Text: &vn.TextNode{Value: "Hello world", Element: nil}, Element: nil}},
		{&vn.Attrs{Events: ev}, &vn.Vnode{TagName: "div", Attrs: &vn.Attrs{Events: ev}, Text: &vn.TextNode{Value: "Hello world", Element: nil}, Element: nil}},
		{&vn.Attrs{Events: ev, Props: &vn.Props{"class": "hello"}}, &vn.Vnode{TagName: "div", Attrs: &vn.Attrs{Events: ev, Props: &vn.Props{"class": "hello"}}, Text: &vn.TextNode{Value: "Hello world", Element: nil}, Element: nil}},
	}
	for _, c := range cases {
		got := vn.H("div", c.Attrs, "Hello world")

		assert.Equal(t, c.want, got)
	}
}
