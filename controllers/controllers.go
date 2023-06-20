package controllers

import (
	"errors"
	"net/http"

	"go-rest-api/data"

	"github.com/gin-gonic/gin"
)

// * Controller to get todos
func GetTodos(c *gin.Context){
	c.IndentedJSON(http.StatusOK, data.Todos)
}

// * Controller to add todo
func AddTodo(c *gin.Context) {
	var newTodo data.Todo;
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err);
		return
	}
	data.Todos = append(data.Todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// * Controller to toggle todo
func ToggleTodo(c *gin.Context){
	id := c.Param("id");
	todo, err := getTodoByID(id);
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"});
		return;
	}
	todo.Completed = !todo.Completed
	c.IndentedJSON(http.StatusOK, todo)
}

// * Controller to get single todo
func GetTodo(c *gin.Context){
	id := c.Param("id");
	todo, err := getTodoByID(id);
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"});
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

// * Handler to loop through todo slice
func getTodoByID(id string) (*data.Todo, error) {
	for i, todo := range data.Todos {
		if todo.ID == id {
			return &data.Todos[i], nil
		}
	}
	return nil, errors.New("Todo not found!")
}
