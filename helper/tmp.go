package helper

import (
	"errors"
	"os"
	"path/filepath"
)

func CleanUpTemp() {
	files, _ := filepath.Glob("/tmp/remusika-tmpimg-*")
	for _, f := range files {
		os.Remove(f)
	}
}

func CreateImageTemp(imageBytes []byte) (string, error) {
	if imageBytes == nil {
		return "", errors.New("image is null")
	}

	tmpFile, err := os.CreateTemp("", "remusika-tmpimg-*")
	if err != nil {
		return "", err
	}

	defer tmpFile.Close()

	_, err = tmpFile.Write(imageBytes)
	if err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}
