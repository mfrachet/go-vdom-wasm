package cmps

import vn "github.com/mfrachet/go-vdom-wasm"

func Notif(text string) *vn.Vnode {
	return vn.H("div", &vn.Attrs{Props: &vn.Props{"class": "notification m-t"}}, text)
}
