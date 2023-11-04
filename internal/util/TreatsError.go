package util

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TreatsError(c *gin.Context, err error, errResponse ErrorResponse, logErrMessage string) {
	errResponse.Message = err.Error()
	errResponse.ErrorCode = http.StatusBadRequest
	log.Printf("%s: %v", logErrMessage, err)
	SendError(c, errResponse)
}
