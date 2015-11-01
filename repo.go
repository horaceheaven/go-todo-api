package main

import (
	"fmt"
)

// Obviously this solution isn't thread safe

var currentId int
var todos Todos


// Seed data
func init() {
	RepoCreateTodo(Todo{Name: "Wash the car"})
	RepoCreateTodo(Todo{Name: "Walk the dogs"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId ++
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not find todo with id of %d to delete", id)
}