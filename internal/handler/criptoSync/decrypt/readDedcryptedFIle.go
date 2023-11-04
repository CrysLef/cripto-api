package decrypt

import (
	"crypto/cipher"
	"io"
	"log"
	"os"
)

func readEncryptedFile(inFile, outFile *os.File, msgLen int64, buf []byte, stream cipher.Stream) error {
	for {
		n, err := inFile.Read(buf)
		if n > 0 {
			if n > int(msgLen) {
				n = int(msgLen)
			}
			msgLen -= int64(n)
			stream.XORKeyStream(buf, buf[:n])
			outFile.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("tamanho de bytes lidos: %d", n)
			return err
		}
	}
	return nil
}
