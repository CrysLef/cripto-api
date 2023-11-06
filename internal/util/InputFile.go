package util

import "os"

func InputFile(file string) (*os.File, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return f, nil
}
