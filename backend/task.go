package main

import "github.com/gin-gonic/gin"

func CreateTask(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Title string `json:"title"`
	}

	c.ShouldBindJSON(&input)

	DB.Exec("INSERT INTO tasks(title,project_id) VALUES(?,?)", input.Title, id)

	c.JSON(200, gin.H{"message": "task created"})
}

func GetTasks(c *gin.Context) {
	id := c.Param("id")

	rows, _ := DB.Query("SELECT id,title,status FROM tasks WHERE project_id=?", id)

	var tasks []gin.H

	for rows.Next() {
		var tid int
		var title, status string

		rows.Scan(&tid, &title, &status)

		tasks = append(tasks, gin.H{
			"id": tid,
			"title": title,
			"status": status,
		})
	}

	c.JSON(200, gin.H{"tasks": tasks})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	DB.Exec("UPDATE tasks SET status='done' WHERE id=?", id)

	c.JSON(200, gin.H{"message": "updated"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	DB.Exec("DELETE FROM tasks WHERE id=?", id)

	c.JSON(200, gin.H{"message": "deleted"})
}

