package main

import (
	"fmt"
	config "todoapp/internal/config"
	controllers "todoapp/internal/http/Controllers"
	db "todoapp/internal/models"

	"github.com/gin-gonic/gin"
)

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

	// controllers
	// todo controllers
	todoControllers := r.Group("/todos")

	todoControllers.GET("/", controllers.GetTodos)
	todoControllers.GET("/:id", controllers.GetTodo)
	todoControllers.POST("/", controllers.CreateTodo)
	todoControllers.PATCH("/:id", controllers.UpdateTodo)

	// run server
	addr := fmt.Sprintf(":%d", env.PORT)
  r.Run(addr)
}