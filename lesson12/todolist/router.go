package todolist

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "PUT", "POST", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Contetn-Length", "Content-Type"}

	router.Use(cors.New(config))

	router.GET("/tasks", handler.GetTaskHandler)
	router.POST("/tasks", handler.CreateTaskHandler)
}
