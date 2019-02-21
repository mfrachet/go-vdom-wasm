package cmps

import (
	"github.com/mfrachet/go-vdom-wasm"
)

func Container(children vn.Children) *vn.Vnode {
	return vn.H("div", &vn.Attrs{Props: &vn.Props{"class": "container grid-lg"}}, children)
}
