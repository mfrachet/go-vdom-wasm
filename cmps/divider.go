package cmps

import (
	"github.com/mfrachet/go-vdom-wasm"
)

func Divider() *vn.Vnode {
	return vn.H("div", &vn.Attrs{Props: &vn.Props{"class": "divider"}}, "")
}
