package main

import (
	vn "github.com/mfrachet/go-vdom-wasm"
	"github.com/mfrachet/vdom-test/cmps"
	"github.com/mfrachet/vdom-test/todo"
)

func main() {
	rootNode := vn.H("div", nil, vn.Children{
		cmps.Navbar(),
		cmps.Container(vn.Children{
			cmps.Title("1", "Built with Go, for Webassembly, with a VDOM"),
			cmps.Divider(),
			todo.TodoList(),
		}),
	})

	vn.Patch("#app", rootNode)
	select {}
}
