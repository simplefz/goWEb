package routes

import (
	"github.com/fanpengjie/project/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterApiRouter(router *gin.Engine)  {
	api := router.Group("/api")
	api.GET("/index/index",controllers.IndexApi)

	api.GET("/import/student/:projectId",controllers.ImportStudent)
}