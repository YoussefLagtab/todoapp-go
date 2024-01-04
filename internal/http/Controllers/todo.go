package controllers

import (
	"net/http"
	"strconv"
	db "todoapp/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateTodoInput struct {
  Content  string `json:"content" binding:"required"`
}
type UpdateTodoInput struct {
  Content  string `json:"content" binding:"required"`
  IsComplete  *bool `json:"isComplete" binding:"required"`
}

func GetTodo(c *gin.Context) {
	var todo db.Todo

	if err := db.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
    return
  }


	c.JSON(http.StatusOK, gin.H{"data": todo})
}
func GetTodos(c *gin.Context) {
	var todos []db.Todo
	db.DB.Order("id ASC").Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func CreateTodo(c *gin.Context) {
	var input CreateTodoInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
	todo := db.Todo{Content: input.Content, IsComplete: false}
	db.DB.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	var input UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	todo := db.Todo{ID: uint(id)}

	db.DB.Model(&todo).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}
