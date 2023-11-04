package decrypt

import (
	"log"

	"github.com/CrysLef/cripto-api/internal/util"
)

func decryptFiles(fileInput, fileOutput string, key []byte) error {
	buf := make([]byte, 4096)

	inFile, err := util.InputFile(fileInput)
	if err != nil {
		return err
	}

	outFile, err := util.OutputFile(fileOutput)
	if err != nil {
		return err
	}

	fi, err := inFile.Stat()
	if err != nil {
		return err
	}

	iv := make([]byte, len(key))

	msgLen := fi.Size() - int64(len(iv))
	if _, err = inFile.ReadAt(iv, msgLen); err != nil {
		log.Fatal("erro ao ler iv: ", err)
	}

	stream, err := util.CreateCTRValidator(key, iv)
	if err != nil {
		return err
	}

	if err := readEncryptedFile(inFile, outFile, msgLen, buf, stream); err != nil {
		return err
	}

	return nil
}
