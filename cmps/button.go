package cmps

import (
	"syscall/js"

	vn "github.com/mfrachet/go-vdom-wasm"
)

func Button(value string, handler func([]js.Value)) *vn.Vnode {
	return vn.H("button", &vn.Attrs{Props: &vn.Props{"class": "button is-primary m-r"}, Events: &vn.Ev{"click": handler}}, value)
}
