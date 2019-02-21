package cmps

import vn "github.com/mfrachet/go-vdom-wasm"

func Avatar(imageLink string) *vn.Vnode {
	return vn.H("figure", &vn.Attrs{Props: &vn.Props{"class": "avatar"}}, vn.Children{
		vn.H("img", &vn.Attrs{Props: &vn.Props{
			"src": imageLink,
		}}, ""),
	})
}
