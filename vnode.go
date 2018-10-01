package vn

import (
	"reflect"
	"syscall/js"
)

type Ev = map[string]func([]js.Value)
type Props = map[string]string
type Attrs struct {
	Props  *Props
	Events *Ev
}

type Vnode struct {
	TagName  string
	Attrs    *Attrs
	Children Node
	Element  *js.Value
}

func (vnode *Vnode) isSame(otherVnode *Vnode) bool {
	if vnode.Attrs == nil && otherVnode.Attrs == nil {
		return vnode.TagName == otherVnode.TagName
	}

	if vnode.Attrs.Props == nil && otherVnode.Attrs.Props == nil {
		return vnode.TagName == otherVnode.TagName
	}

	return vnode.TagName == otherVnode.TagName && reflect.DeepEqual(vnode.Attrs.Props, otherVnode.Attrs.Props)
}

func (vnode *Vnode) createElement() {
	document := js.Global().Get("document")
	domNode := document.Call("createElement", vnode.TagName)

	if vnode.Attrs != nil {
		if vnode.Attrs.Props != nil {
			for k, v := range *vnode.Attrs.Props {
				domNode.Call("setAttribute", k, v)
			}
		}

		if vnode.Attrs.Events != nil {
			for k, f := range *vnode.Attrs.Events {
				callback := js.NewCallback(f)
				domNode.Call("addEventListener", k, callback)
			}
		}

	}

	vnode.Children.createElement()
}
