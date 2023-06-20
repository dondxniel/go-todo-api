package main

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default();
	router.GET("/todos", controllers.GetTodos)
	router.GET("/todos/:id", controllers.GetTodo)
	router.PATCH("/toggle-todo/:id", controllers.ToggleTodo)
	router.POST("/todos", controllers.AddTodo)
	router.Run("localhost:8000")
}