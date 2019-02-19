package vn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameVNode(t *testing.T) {
	a := H("div", nil, "Hello world")
	b := H("div", nil, "Hello world")
	c := H("span", nil, "Hello world")

	d := H("span", &Attrs{Props: &Props{"hello": "world"}}, "Hello world")
	e := H("span", &Attrs{Props: &Props{"hello": "world"}}, "Hello world")
	f := H("div", &Attrs{Props: &Props{"hello": "world"}}, "Hello worldx")
	j := H("span", &Attrs{Props: &Props{"hello": "world2"}}, "Hello world")

	g := H("span", &Attrs{}, "Hello world")
	h := H("span", &Attrs{}, "Hello world")

	k := H("li", nil, "Hello world")
	l := H("li", nil, "Hello world2")

	assert.Equal(t, true, a.IsSame(b))
	assert.Equal(t, false, a.IsSame(c))
	assert.Equal(t, true, g.IsSame(h))
	assert.Equal(t, true, d.IsSame(e))
	assert.Equal(t, false, d.IsSame(f))
	assert.Equal(t, false, d.IsSame(j))
	assert.Equal(t, false, k.IsSame(l))
}

func TestVnodeChildrenCount(t *testing.T) {
	a := H("div", nil, Children{
		H("span", nil, "Hello"),
		H("span", nil, "Hello"),
	})

	assert.Equal(t, 2, a.ChildrenCount())
}

// func TestVnodeCreateElement(t *testing.T) {
// 	document := js.Global().Get("document")
// 	domNode := document.Call("createElement", "div")
// 	spanNode := document.Call("createElement", "span")
// 	textNode := document.Call("createTextNode", "Hello")
// 	spanNode.Call("appendChild", textNode)
// 	domNode.Call("appendChild", spanNode)

// 	vnode := H("div", nil, Children{
// 		H("span", nil, "Hello"),
// 	})

// 	vnode.CreateElement()

// 	assert.Equal(t, domNode, *vnode.Element)
// }
