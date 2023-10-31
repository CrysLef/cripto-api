package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	basePath := "/api/v1/sync"

	v1 := router.Group(basePath)
	{
		v1.POST("/encrypt")
		v1.POST("/decrypt")
	}
}
