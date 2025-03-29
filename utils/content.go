package utils

import (
	"fmt"
	"io"
	"mime/multipart"
)

func GetFileContent(file *multipart.FileHeader) (string, error) {
	openedFile, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer openedFile.Close()

	content, err := io.ReadAll(openedFile)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(content), nil
}

