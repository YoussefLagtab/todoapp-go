package main

import (
	"fmt"
	config "todoapp/internal/config"
	controllers "todoapp/internal/http/Controllers"
	db "todoapp/internal/models"

	"github.com/gin-gonic/gin"
)

var todoController controllers.TodoController

func main() {
	// read env
	env := config.ReadEnv()

	// create server
	r := gin.New()

	// connect db
  db.ConnectDatabase(env)

	// run migration if RUN_AUTO_MIGRATION env var is set
	if (env.RUN_AUTO_MIGRATION) {
		db.RunAutoMigartion()
	}

	// routes
	// todo routes
	todoRoutes := r.Group("/todos")

	todoRoutes.GET("/", todoController.GetAllTodos)
	todoRoutes.GET("/:id", todoController.GetTodo)
	todoRoutes.POST("/", todoController.CreateTodo)
	todoRoutes.PATCH("/:id", todoController.UpdateTodo)
	todoRoutes.PATCH("/:id/mark-as-complete", todoController.MarkTodoAsCompleted)
	todoRoutes.PATCH("/:id/mark-as-incomplete", todoController.MarkTodoAsInCompleted)

	// run server
	addr := fmt.Sprintf(":%d", env.PORT)
  r.Run(addr)
}