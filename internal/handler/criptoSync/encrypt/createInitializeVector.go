package encrypt

import (
	"crypto/rand"
	"io"
)

func createInitializeVector(size int) ([]byte, error) {
	iv := make([]byte, size)

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	return iv, nil
}
