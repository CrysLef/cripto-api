package util

import "os"

func OutputFile(file string) (*os.File, error) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	return f, nil
}
