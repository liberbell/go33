package todolist

import "github.com/gin-gonic/gin"

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

func (handler *Handler) UpdateTaskHandlr(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedTask := handler.DB.UptateTask(task)
	c.JSON(200, updatedTask)
}
