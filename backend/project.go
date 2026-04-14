package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProjectInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateProject(c *gin.Context) {
	var input CreateProjectInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userID, _ := c.Get("user_id")

	var id string
	err := DB.QueryRow(
		"INSERT INTO projects (name,description,owner_id) VALUES ($1,$2,$3) RETURNING id",
		input.Name, input.Description, userID,
	).Scan(&id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func GetProjects(c *gin.Context) {
	userID, _ := c.Get("user_id")

	rows, err := DB.Query("SELECT id,name,description FROM projects WHERE owner_id=$1", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}
	defer rows.Close()

	var projects []gin.H

	for rows.Next() {
		var id, name, desc string
		rows.Scan(&id, &name, &desc)

		projects = append(projects, gin.H{
			"id":          id,
			"name":        name,
			"description": desc,
		})
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}