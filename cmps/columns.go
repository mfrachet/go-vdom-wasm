package cmps

import (
	"github.com/mfrachet/go-vdom-wasm"
)

func Columns(children vn.Children) *vn.Vnode {
	return vn.H("div", &vn.Attrs{Props: &vn.Props{"class": "columns"}}, children)
}
