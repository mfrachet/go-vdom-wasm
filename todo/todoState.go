package todo

type TodoState struct {
	todos []*TodoModel
}

func indexOf(slice []*TodoModel, todo *TodoModel) int {
	size := len(slice)

	for i := 0; i < size; i++ {
		if slice[i].text == todo.text {
			return i
		}
	}

	return -1
}

func (todoState *TodoState) Count() int {
	return len(todoState.todos)
}

func (todoState *TodoState) ChildAt(index int) *TodoModel {
	return todoState.todos[index]
}

func (todoState *TodoState) AddTodo(value string) {
	todo := &TodoModel{value, false}

	indexOfExisting := indexOf(todoState.todos, todo)
	if indexOfExisting == -1 {
		todoState.todos = append(todoState.todos, todo)
	}
}

func (todoState *TodoState) RemoveTodo(todo *TodoModel) {
	i := indexOf(todoState.todos, todo)

	todoState.todos = append(todoState.todos[:i], todoState.todos[i+1:]...)
}

func (todoState *TodoState) UpdateTodo(todo *TodoModel) {
	todo.isChecked = !todo.isChecked
}
