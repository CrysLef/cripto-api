package encrypt

import (
	"crypto/cipher"
	"io"
	"os"
)

func writeEncryptedFile(inFile, outFile *os.File, buf []byte, stream cipher.Stream) error {
	for {
		n, err := inFile.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			outFile.Write(buf[:n])
		}
		if err != io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}
