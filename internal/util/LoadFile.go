package util

import (
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoadFile(c *gin.Context, errResponse ErrorResponse) (string, string, []byte, error) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		TreatsError(c, err, errResponse, "erro ao ler o arquivo")
		return "", "", nil, err
	}

	fileName := uuid.New().String()
	tempFileName := uuid.New().String()
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		TreatsError(c, err, errResponse, "erro ao ler o arquivo")
		return "", "", nil, err
	}
	log.Println("arquivo carregado com sucesso")

	return fileName, tempFileName, fileBytes, nil
}
