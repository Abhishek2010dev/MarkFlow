package utils

import (
	"fmt"
	"io"
	"mime/multipart"
)

func GetFileContent(file *multipart.File) (string, error) {
	content, err := io.ReadAll(*file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(content), nil
}

