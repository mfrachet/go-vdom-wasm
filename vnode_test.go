package vn

import (
	"testing"

	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	"github.com/stretchr/testify/assert"
)

func TestVnode_IsSame(t *testing.T) {
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
	assert.Equal(t, true, g.IsSame(h))
	assert.Equal(t, true, d.IsSame(e))

	assert.Equal(t, false, a.IsSame(c))
	assert.Equal(t, false, d.IsSame(f))
	assert.Equal(t, false, d.IsSame(j))
	assert.Equal(t, false, k.IsSame(l))
}

func TestVnode_ChildrenCount(t *testing.T) {
	a := H("div", nil, Children{
		H("span", nil, "Hello"),
		H("span", nil, "Hello"),
	})

	assert.Equal(t, 2, a.ChildrenCount())
}

func TestVnode_ChildAt(t *testing.T) {
	childAtOne := H("span", nil, "Hello World")

	a := H("div", nil, Children{
		H("span", nil, "Hello"),
		childAtOne,
	})
	b := H("span", nil, "Hello")

	assert.Equal(t, childAtOne, a.ChildAt(1))
	assert.Equal(t, nil, b.ChildAt(1))
}

func TestVnode_HasElement(t *testing.T) {
	element := H("span", nil, "Hello world")
	assert.Equal(t, false, element.HasElement())

	element.SetElement(vnd.DomElement{})
	assert.Equal(t, true, element.HasElement())
}
