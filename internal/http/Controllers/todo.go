package controllers

import (
	"net/http"
	helpers "todoapp/internal"
	services "todoapp/internal/services/todo"

	"github.com/gin-gonic/gin"
)

type TodoController struct{}

type createTodoInput struct {
	Content string `json:"content" binding:"required"`
}
type updateTodoContentInput struct {
	Content    string `json:"content" binding:"required"`
}

var todoService services.TodoService;

func (tc *TodoController) GetTodo(c *gin.Context) {
	id, err := helpers.GetIdFromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	todo, err := todoService.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func (tc *TodoController) GetAllTodos(c *gin.Context) {
	todos := todoService.GetAllTodos()

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func (tc *TodoController) CreateTodo(c *gin.Context) {
	var input createTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := todoService.CreateTodo(input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func (tc *TodoController) UpdateTodo(c *gin.Context) {
	var input updateTodoContentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := helpers.GetIdFromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	todo, err := todoService.UpdateTodoContent(id, input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func (tc *TodoController) MarkTodoAsCompleted(c *gin.Context) {
	id, err := helpers.GetIdFromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	todo, err := todoService.MarkTodoAsCompleted(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func (tc *TodoController) MarkTodoAsInCompleted(c *gin.Context) {
	id, err := helpers.GetIdFromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	todo, err := todoService.MarkTodoAsInCompleted(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}
