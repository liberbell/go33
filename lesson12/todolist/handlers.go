package todolist

import "github.com/gin-gonic/gin"

type Handler struct {
	Engine *gin.Engine
	DB     *Database
}

func (handler *Handler) GetTasksHandler(c *gin.Context) {

}
