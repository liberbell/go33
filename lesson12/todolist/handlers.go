package todolist

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine *gin.Engine
	DB     *Database
}

func (handler *Handler) GetTasksHandler(c *gin.Context) {
	tasks := handler.DB.GetTasks()
	c.JSON(200, tasks)
}

func (handler *Handler) CreateTaskHandler(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createTask := handler.DB.CreateTask(task.Title, task.Description, task.Completed)
	c.JSON(201, createTask)
}

func (handler *Handler) UpdateTaskHandler(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedTask := handler.DB.UpdateTask(task)
	c.JSON(200, updatedTask)
}

func (handler *Handler) DeleteTaskHandler(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	id := uint(idInt)
	success := handler.DB.DeleteTask(id)
	if success {
		c.JSON(400, gin.H{"error": "Failed to delete task"})
	}
}
