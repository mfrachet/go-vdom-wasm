package cmps

import (
	"github.com/mfrachet/go-vdom-wasm"
)

func Menu(children vn.Children) *vn.Vnode {
	return vn.H("ul", &vn.Attrs{Props: &vn.Props{"class": "nav"}}, children)
}
