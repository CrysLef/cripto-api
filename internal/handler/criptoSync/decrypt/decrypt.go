package decrypt

import (
	"log"
	"net/http"
	"os"

	"github.com/CrysLef/cripto-api/internal/util"
	"github.com/gin-gonic/gin"
)

func Run(c *gin.Context) {
	w := c.Writer
	errResponse := util.ErrorResponse{}

	fileName, tempFileName, fileBytes, err := util.LoadFile(c, errResponse)
	if err != nil {
		errResponse.Message = err.Error()
		errResponse.ErrorCode = http.StatusBadRequest
		log.Printf("erro ao carregar o arquivo: %v", err)
		util.SendError(c, errResponse)
		return
	}

	tempFile := util.CreateTempFile(c, fileName, fileBytes, errResponse)

	if err := decryptFiles(tempFile.Name(), tempFileName, []byte("123456789qwertyu")); err != nil {
		log.Fatalln(err)
	}
	log.Println("arquivo criptografado com sucesso")

	returnFileBytes, err := os.ReadFile(tempFileName)
	if err != nil {
		util.TreatsError(c, err, errResponse, "erro ao ler arquivo descriptografado")
		return
	}

	util.SendSuccess(w, c, http.StatusOK, returnFileBytes)

	defer os.Remove(tempFileName)
	defer os.Remove(tempFile.Name())

}
