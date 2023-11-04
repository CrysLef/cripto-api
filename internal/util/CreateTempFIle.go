package util

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateTempFile(c *gin.Context, fileName string, fileBytes []byte, errResponse ErrorResponse) *os.File {
	tempFile, err := os.CreateTemp("./", fileName)
	if err != nil {
		TreatsError(c, err, errResponse, "erro criando o arquivo")
		return nil
	}
	defer tempFile.Close()
	tempFile.Write(fileBytes)
	log.Println("arquivo temporario criado com sucesso")

	return tempFile
}
