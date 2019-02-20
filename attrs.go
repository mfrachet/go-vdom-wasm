package vn

import (
	"syscall/js"

	vnh "github.com/mfrachet/go-vdom-wasm/helpers"
)

type Ev = map[string]func([]js.Value)

type Props = map[string]string

type Attrs struct {
	Props  *Props
	Events *Ev
}

func Sanitize(attrs *Attrs) Attrs {
	if vnh.IsNil(attrs) {
		return Attrs{Props: &Props{}, Events: &Ev{}}
	}

	if vnh.IsNil(attrs.Props) {
		attrs.Props = &Props{}
	}

	if vnh.IsNil(attrs.Events) {
		attrs.Events = &Ev{}
	}

	return *attrs
}
