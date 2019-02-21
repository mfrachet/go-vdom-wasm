package cmps

import vn "github.com/mfrachet/go-vdom-wasm"

func Title(size string, value string) *vn.Vnode {
	return vn.H("h"+size, &vn.Attrs{Props: &vn.Props{"class": "title is-" + size}}, value)
}
