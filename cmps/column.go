package cmps

import (
	"github.com/mfrachet/go-vdom-wasm"
)

func Column(size string, isColored bool, children vn.Children) *vn.Vnode {
	classes := "column col-" + size
	if isColored {
		classes += " bg-secondary"
	}

	return vn.H("div", &vn.Attrs{Props: &vn.Props{"class": classes}}, children)
}
