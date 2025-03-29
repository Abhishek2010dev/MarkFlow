package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func FileExits(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func GetFileContent(file *multipart.File) ([]byte, error) {
	content, err := io.ReadAll(*file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return content, nil
}

func ReadFileContent(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filename, err)
	}
	return content, nil
}
