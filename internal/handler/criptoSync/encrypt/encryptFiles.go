package encrypt

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/CrysLef/cripto-api/internal/util"
)

func encryptFiles(fileInput, fileOutput string, key []byte) error {
	buf := make([]byte, 4096)

	inFile, err := util.InputFile(fileInput)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := util.OutputFile(fileOutput)
	if err != nil {
		return err
	}
	defer outFile.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv, err := createInitializeVector(block.BlockSize())
	if err != nil {
		return err
	}

	stream := cipher.NewCTR(block, iv)

	if err := writeEncryptedFile(inFile, outFile, buf, stream); err != nil {
		return err
	}

	outFile.Write(iv)
	return nil
}
