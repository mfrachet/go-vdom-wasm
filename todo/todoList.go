package todo

import (
	"syscall/js"
	"time"

	vn "github.com/mfrachet/go-vdom-wasm"
	"github.com/mfrachet/vdom-test/cmps"
)

var inputValue string
var rows *vn.Vnode
var delay *vn.Vnode
var todos *TodoState

func apply() {
	start := time.Now()
	(func() {
		size := todos.Count()
		nextChild := vn.Children{}

		for i := 0; i < size; i++ {
			nextChild = append(nextChild, createItem(todos.ChildAt(i)))
		}

		nextList := createList(nextChild)
		vn.Patch(rows, nextList)
		rows = nextList
	}())

	t := time.Now()
	elapsed := t.Sub(start)

	nextNotif := createNotif("It has taken " + elapsed.String())
	vn.Patch(delay, nextNotif)

	delay = nextNotif
}

func createList(children vn.Children) *vn.Vnode {
	return vn.H("nav", &vn.Attrs{Props: &vn.Props{"class": "panel"}}, children)
}

func createItem(todo *TodoModel) *vn.Vnode {
	var dynamicClass string

	if todo.isChecked {
		dynamicClass = " is-checked"
	}

	return vn.H("a", &vn.Attrs{Props: &vn.Props{"class": "panel-block" + dynamicClass}}, vn.Children{
		vn.H("span", nil, todo.text),
		vn.H("button", &vn.Attrs{Props: &vn.Props{"class": "button is-primary m-l"}, Events: &vn.Ev{"click": handleCheck(todo)}}, "Check me :D"),
		vn.H("button", &vn.Attrs{Props: &vn.Props{"class": "button is-danger m-l"}, Events: &vn.Ev{"click": handleRemove(todo)}}, "Delete me :("),
	}, todo.text)
}

func createNotif(notif string) *vn.Vnode {
	return cmps.Notif(notif)
}

func handleKeyUp(values []js.Value) {
	keyCode := values[0].Get("keyCode").Int()
	value := values[0].Get("target").Get("value").String()

	inputValue = value

	if keyCode == 13 {
		todos.AddTodo(inputValue)
		apply()

	}
}

func handleCheck(todo *TodoModel) func([]js.Value) {
	return func(args []js.Value) {
		todos.UpdateTodo(todo)
		apply()
	}
}

func handleRemove(todo *TodoModel) func([]js.Value) {
	return func(args []js.Value) {
		todos.RemoveTodo(todo)
		apply()
	}
}

func TodoList() *vn.Vnode {
	todos = &TodoState{[]*TodoModel{}}

	rows = createList(vn.Children{})
	delay = createNotif("Make an action on the todo to check for the time it takes")

	return vn.H("div", nil, vn.Children{
		cmps.Input(handleKeyUp),
		delay,
		rows,
	})
}
