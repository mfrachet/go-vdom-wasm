package vn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

import vn "github.com/mfrachet/go-vdom-wasm"

func TestH_WithText(t *testing.T) {
	expected := vn.NewNode("div", &vn.Attrs{Props: &vn.Props{}, Events: &vn.Ev{}}, nil, &vn.TextNode{Value: "Hello world", Element: nil}, nil, nil)

	vNode := vn.H("div", nil, "Hello world")

	assert.Equal(t, expected, vNode)
}

func TestH_WithChildren(t *testing.T) {
	child := vn.H("span", nil, "Hello world")
	expected := vn.NewNode("div", &vn.Attrs{Props: &vn.Props{}, Events: &vn.Ev{}}, vn.Children{child}, nil, nil, nil)

	vNode := vn.H("div", nil, vn.Children{child})

	assert.Equal(t, expected, vNode)
}

func TestH_WithAttributes(t *testing.T) {
	exectedWithoutAttrs := vn.NewNode("div", &vn.Attrs{Props: &vn.Props{}, Events: &vn.Ev{}}, nil, &vn.TextNode{Value: "Hello world", Element: nil}, nil, nil)
	vNodeWithoutAttrs := vn.H("div", nil, "Hello world")

	assert.Equal(t, exectedWithoutAttrs, vNodeWithoutAttrs)

	exectedWithAttrs := vn.NewNode("div", &vn.Attrs{Props: &vn.Props{"class": "navbar"}, Events: &vn.Ev{}}, nil, &vn.TextNode{Value: "Hello world", Element: nil}, nil, nil)
	vNodeWithAttrs := vn.H("div", &vn.Attrs{Props: &vn.Props{"class": "navbar"}}, "Hello world")

	assert.Equal(t, exectedWithAttrs, vNodeWithAttrs)
}
