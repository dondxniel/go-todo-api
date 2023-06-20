package controllers

import (
	"net/http"

	"go-rest-api/constants"
	"go-rest-api/data"
	"go-rest-api/services"

	"github.com/gin-gonic/gin"
)

// * Controller to get todos
func GetTodos(c *gin.Context){
	c.IndentedJSON(http.StatusOK, services.FetchTodosService())
}

// * Controller to add todo
func AddTodo(c *gin.Context) {
	var newTodo data.Todo;
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err);
		return
	}
	services.AddTodoService(newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// * Controller to toggle todo
func ToggleTodo(c *gin.Context){
	id := c.Param("id");
	todo, err := services.GetSingleTodo(id);
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": constants.Todo404Msg});
		return;
	}
	todo.Completed = !todo.Completed
	c.IndentedJSON(http.StatusOK, todo)
}

// * Controller to get single todo
func GetTodo(c *gin.Context){
	id := c.Param("id");
	todo, err := services.GetSingleTodo(id);
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": constants.Todo404Msg});
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}
