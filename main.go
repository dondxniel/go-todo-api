package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID string `json:"id"`
	Item string  `json:"item"`
	Completed bool `json:"completed"`
}

var todos = []todo{
	{ ID: "1", Item: "Wash", Completed: false },
	{ ID: "2", Item: "Code", Completed: false },
	{ ID: "3", Item: "Read", Completed: false },
}

// * Controller to get todo
func getTodos(c *gin.Context){
	c.IndentedJSON(http.StatusOK, todos)
}

// * Controller to add todo
func addTodo(c *gin.Context) {
	var newTodo todo;
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err);
		return
	}
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// * Controller to toggle todo item status
func toggleTodo(c *gin.Context) {
	id := c.Param("id");
	todo, err := getTodoByID(id);
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"});
		return;
	}
	todo.Completed = !todo.Completed
	c.IndentedJSON(http.StatusOK, todo)
}

// * Controller to get todo by ID
func getTodo(c *gin.Context){
	id := c.Param("id");
	todo, err := getTodoByID(id);
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"});
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

// * Handler to loop through todo slice
func getTodoByID(id string) (*todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found!")
}

func main(){
	router := gin.Default();
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/toggle-todo/:id", toggleTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:8000")
}