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

type Children []*Vnode

type Vnode struct {
	TagName  string
	Attrs    *Attrs
	Children interface{}
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
