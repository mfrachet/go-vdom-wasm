package cmps

import (
	vn "github.com/mfrachet/go-vdom-wasm"
)

func Navbar() *vn.Vnode {
	IMAGE_ADRESS := "https://dashboard.snapcraft.io/site_media/appmedia/2017/11/gophercolor.png"

	brandSection := vn.H("div", &vn.Attrs{Props: &vn.Props{"class": "navbar-brand"}}, vn.Children{
		vn.H("a", &vn.Attrs{Props: &vn.Props{"class": "navbar-item"}}, vn.Children{Avatar(IMAGE_ADRESS)}),
		vn.H("span", &vn.Attrs{Props: &vn.Props{"class": "navbar-item"}}, vn.Children{Title("6", "go-vdom-wasm")}),
	})

	menuContent := vn.H("div", &vn.Attrs{Props: &vn.Props{"class": "navbar-end"}}, vn.Children{
		vn.H("a", &vn.Attrs{Props: &vn.Props{"class": "navbar-item", "href": "https://github.com/mfrachet/go-vdom-wasm", "target": "_blank"}}, "Github"),
		vn.H("a", &vn.Attrs{Props: &vn.Props{"class": "navbar-item", "href": "https://twitter.com/mfrachet", "target": "_blank"}}, "Twitter"),
		vn.H("a", &vn.Attrs{Props: &vn.Props{"class": "navbar-item", "href": "https://medium.com/@mfrachet", "target": "_blank"}}, "Medium"),
	})

	return vn.H("nav", &vn.Attrs{Props: &vn.Props{"class": "navbar is-primary"}}, vn.Children{
		brandSection,
		menuContent,
	})
}
