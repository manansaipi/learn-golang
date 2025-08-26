package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json: "completed"`
}

var todos = []Todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Work", Completed: true},
	{ID: "3", Item: "Walking", Completed: false},
}

func getTodos(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, todos)
}
func addTodos(context *gin.Context) {
	var newTodo Todo
	if err := context.BindJSON(&newTodo); err != nil {
		// Respond with a Bad Request error if JSON binding fails
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)

}

func getTodoById(id string) (*Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	Todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, Todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	Todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	Todo.Completed = !Todo.Completed

	context.IndentedJSON(http.StatusOK, Todo)

}

func main() {

	router := gin.Default()
	router.GET("/", defaultHandler)
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodos)
	router.Run("localhost:9090")

}
func defaultHandler(c *gin.Context) {
	c.String(200, "connection success")
}
