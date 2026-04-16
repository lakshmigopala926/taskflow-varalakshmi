package main

import (
	"fmt"
	"time"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	InitDB()
	CreateTables()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/login", Login)

	// public (for simplicity)
	r.POST("/projects", CreateProject)
	r.GET("/projects", GetProjects)
	r.DELETE("/projects/:id", DeleteProject)

	r.POST("/projects/:id/tasks", CreateTask)
	r.GET("/projects/:id/tasks", GetTasks)
	r.PATCH("/tasks/:id", UpdateTask)
	r.DELETE("/tasks/:id", DeleteTask)

	port := os.Getenv("PORT")
if port == "" {
	port = "8080"
}

r.Run(":" + port)
}