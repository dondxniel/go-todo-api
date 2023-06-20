package main

import (
	"go-rest-api/constants"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default();
	router.GET(constants.TodosRoute, controllers.GetTodos)
	router.GET(constants.SingleTodoRoute, controllers.GetTodo)
	router.PATCH(constants.ToggleTodoRoute, controllers.ToggleTodo)
	router.POST(constants.TodosRoute, controllers.AddTodo)
	router.Run("localhost:8000")
}