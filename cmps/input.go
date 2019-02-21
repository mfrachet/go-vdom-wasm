package cmps

import (
	"syscall/js"

	vn "github.com/mfrachet/go-vdom-wasm"
)

func Input(handler func([]js.Value)) *vn.Vnode {
	return vn.H("div", &vn.Attrs{Props: &vn.Props{"class": "field"}}, vn.Children{
		vn.H("input", &vn.Attrs{Props: &vn.Props{"type": "text", "class": "input m-t", "placeholder": "Type something here :)"}, Events: &vn.Ev{"keyup": handler}}, ""),
	})
}
