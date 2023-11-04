package encrypt

import "github.com/CrysLef/cripto-api/internal/util"

func encryptFiles(fileInput, fileOutput string, key []byte) error {
	buf := make([]byte, 4096)

	inFile, err := util.InputFile(fileInput)
	if err != nil {
		return err
	}

	outFile, err := util.OutputFile(fileOutput)
	if err != nil {
		return err
	}

	iv, err := createInitializeVector(len(key))
	if err != nil {
		return err
	}

	stream, err := util.CreateCTRValidator(key, iv)
	if err != nil {
		return nil
	}

	if err := writeEncryptedFile(inFile, outFile, buf, stream); err != nil {
		return err
	}

	outFile.Write(iv)
	return nil
}
