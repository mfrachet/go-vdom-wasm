package cmps

import (
	"github.com/mfrachet/go-vdom-wasm"
)

func MenuItem(textElement string, isActive bool) *vn.Vnode {
	classes := "nav-item"

	if isActive {
		classes += " active"
	}

	return vn.H("li", &vn.Attrs{Props: &vn.Props{"class": classes}}, vn.Children{
		vn.H("a", nil, textElement),
	})
}
