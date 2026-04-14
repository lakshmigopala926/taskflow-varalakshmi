package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

func CreateTask(c *gin.Context) {
	projectID := c.Param("id")

	var input CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	var id string
	err := DB.QueryRow(
		"INSERT INTO tasks (title,description,priority,project_id) VALUES ($1,$2,$3,$4) RETURNING id",
		input.Title, input.Description, input.Priority, projectID,
	).Scan(&id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func GetTasks(c *gin.Context) {
	projectID := c.Param("id")

	rows, err := DB.Query("SELECT id,title,status FROM tasks WHERE project_id=$1", projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed"})
		return
	}
	defer rows.Close()

	var tasks []gin.H

	for rows.Next() {
		var id, title, status string
		rows.Scan(&id, &title, &status)

		tasks = append(tasks, gin.H{
			"id":     id,
			"title":  title,
			"status": status,
		})
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var input map[string]interface{}
	c.ShouldBindJSON(&input)

	_, err := DB.Exec("UPDATE tasks SET status=$1 WHERE id=$2", input["status"], id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	_, err := DB.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}

	c.Status(http.StatusNoContent)
}