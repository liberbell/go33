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
}
