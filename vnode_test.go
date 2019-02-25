package vn

import (
	"testing"

	vnd "github.com/mfrachet/go-vdom-wasm/dom"
	"github.com/stretchr/testify/assert"
)

func TestVnode_IsSame(t *testing.T) {
	a := H("div", "Hello world")
	b := H("div", "Hello world")
	c := H("span", "Hello world")

	d := H("span", &Props{"hello": "world"}, "Hello world")
	e := H("span", &Props{"hello": "world"}, "Hello world")
	f := H("div", &Props{"hello": "world"}, "Hello worldx")
	j := H("span", &Props{"hello": "world2"}, "Hello world")

	k := H("li", "Hello world")
	l := H("li", "Hello world2")

	m := H("li", Children{})

	assert.Equal(t, true, a.IsSame(b))
	assert.Equal(t, true, d.IsSame(e))

	assert.Equal(t, false, a.IsSame(c))
	assert.Equal(t, false, d.IsSame(f))
	assert.Equal(t, false, d.IsSame(j))
	assert.Equal(t, false, k.IsSame(l))

	assert.Equal(t, false, m.IsSame(a))
	assert.Equal(t, false, a.IsSame(m))
}

func TestVnode_IsSame_WithChildren(t *testing.T) {
	a := H("ul", &Props{"class": "navbar"}, Children{
		H("li", "First item"),
		H("li", "Second item"),
	})

	b := H("ul", &Props{"class": "navbar"}, Children{
		H("li", "First item"),
		H("li", "Second item"),
	})

	assert.Equal(t, true, a.IsSame(b))
}

func TestVnode_IsSame_WithKey(t *testing.T) {
	a := H("ul", &Props{"class": "navbar"}, Children{
		H("li", "First item"),
		H("li", "Second item"),
	}, &Key{"1"})

	b := H("ul", &Props{"class": "navbar"}, Children{
		H("li", "First item"),
		H("li", "Second item"),
	}, &Key{"2"})

	assert.Equal(t, false, a.IsSame(b))
}

func TestVnode_ChildrenCount(t *testing.T) {
	a := H("div", Children{
		H("span", "Hello"),
		H("span", "Hello"),
	})

	assert.Equal(t, 2, a.ChildrenCount())
}

func TestVnode_ChildAt(t *testing.T) {
	childAtOne := H("span", "Hello World")

	a := H("div", Children{
		H("span", "Hello"),
		childAtOne,
	})
	b := H("span", "Hello")

	assert.Equal(t, childAtOne, a.ChildAt(1))
	assert.Equal(t, nil, b.ChildAt(1))
}

func TestVnode_HasElement(t *testing.T) {
	element := H("span", "Hello world")
	assert.Equal(t, false, element.HasElement())

	element.SetElement(vnd.DomElement{})
	assert.Equal(t, true, element.HasElement())
}

func TestVnode_GetTagName(t *testing.T) {
	element := H("span", "Hello world")

	assert.Equal(t, "span", element.GetTagName())
}

func TestVnode_GetAttrs(t *testing.T) {
	props := &Props{"class": "navbar"}
	ev := &Ev{}
	attrs := &Attrs{Props: props, Events: ev}
	element := H("span", props, ev, "Hello world")

	assert.Equal(t, attrs, element.GetAttrs())
}

func TestVnode_GetChildren(t *testing.T) {
	children := Children{}
	element := H("span", children)

	assert.Equal(t, children, element.GetChildren())
}
