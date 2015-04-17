package todoPkg

import "fmt"

var currentId int
var todos Todos

//give us some seed data

func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	//return empty todo
	return Todo{}
}

func RepoDestroyTodo(id int) error {
	for k, t := range todos {
		if t.Id == id {
			todos = append(todos[:k], todos[k+1:]...)
			return nil
		}
	}

	//return error
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
