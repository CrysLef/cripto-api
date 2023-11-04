package util

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

func SendError(c *gin.Context, errResponse ErrorResponse) {
	c.Header("Content-type", "application/json")
	c.JSON(errResponse.ErrorCode, errResponse)
}

func SendSuccess(w gin.ResponseWriter, c *gin.Context, statusCode int, fileBytes []byte) {
	w.WriteHeader(statusCode)
	c.Header("Content-type", "application/octet-stream")
	w.Write(fileBytes)
}
