package services

import (
	"errors"
	"go-rest-api/data"
)

func FetchTodosService() []data.Todo {
	return data.Todos
}

func AddTodoService(todo data.Todo){
	data.Todos = append(data.Todos, todo)
}

func GetSingleTodo (id string) (*data.Todo, error){
	for i, todo := range data.Todos {
		if todo.ID == id {
			return &data.Todos[i], nil
		}
	}
	return nil, errors.New("Todo not found!")
}