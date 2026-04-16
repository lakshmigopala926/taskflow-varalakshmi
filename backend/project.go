package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid"})
		return
	}

	if input.Name == "" {
		c.JSON(400, gin.H{"error": "name required"})
		return
	}

	_, err := DB.Exec("INSERT INTO projects (name) VALUES (?)", input.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "project created"})
}

func GetProjects(c *gin.Context) {
	rows, _ := DB.Query("SELECT id, name FROM projects")
	defer rows.Close()

	var projects []gin.H

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)

		projects = append(projects, gin.H{
			"id":   id,
			"name": name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}
func DeleteProject(c *gin.Context) {
	id := c.Param("id")

	_, err := DB.Exec("DELETE FROM projects WHERE id=?", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "project deleted"})
}

