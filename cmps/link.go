package cmps

import (
	"syscall/js"

	"github.com/mfrachet/go-vdom-wasm"
)

func Link(text string, clickHandler func(event []js.Value)) *vn.Vnode {
	return vn.H("div", &vn.Attrs{Events: &vn.Ev{"click": clickHandler}}, vn.Children{
		vn.H("a", &vn.Attrs{Props: &vn.Props{"class": "btn btn-link"}}, text),
	})
}
