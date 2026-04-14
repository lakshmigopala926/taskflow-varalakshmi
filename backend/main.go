package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting app...")

	InitDB()
	CreateTables()

	r := gin.Default()

	// Public routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/auth/register", Register)
	r.POST("/auth/login", Login)

	// Protected routes
	protected := r.Group("/")
	protected.Use(AuthMiddleware())

	protected.GET("/protected", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		c.JSON(200, gin.H{
			"message": "Protected route",
			"user_id": userID,
		})
	})

	protected.POST("/projects", CreateProject)
	protected.GET("/projects", GetProjects)

	protected.POST("/projects/:id/tasks", CreateTask)
	protected.GET("/projects/:id/tasks", GetTasks)
	protected.PATCH("/tasks/:id", UpdateTask)
	protected.DELETE("/tasks/:id", DeleteTask)

	fmt.Println("Server starting on port 8080...")

	r.Run(":8080")
}