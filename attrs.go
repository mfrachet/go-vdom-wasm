package vn

import "syscall/js"

type Ev = map[string]func([]js.Value)

type Props = map[string]string

type Attrs struct {
	Props  *Props
	Events *Ev
}
