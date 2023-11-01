package routes

import (
	"github.com/CrysLef/cripto-api/internal/handler/criptoSync/decrypt"
	"github.com/CrysLef/cripto-api/internal/handler/criptoSync/encrypt"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	baseURL := router.Group("/api/v1")

	criptoSync := baseURL.Group("cripto-sync")
	{
		criptoSync.POST("/encrypt", encrypt.Run)
		criptoSync.POST("/decrypt", decrypt.Run)
	}
	return router
}
