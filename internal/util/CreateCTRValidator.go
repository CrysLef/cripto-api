package util

import (
	"crypto/aes"
	"crypto/cipher"
)

func CreateCTRValidator(key, iv []byte) (cipher.Stream, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, iv)

	return stream, nil
}
