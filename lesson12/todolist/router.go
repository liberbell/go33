package todolist

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
}
