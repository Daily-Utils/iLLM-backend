package utils

import (
	"bufio"
	"mime/multipart"
)

func GetTextFromFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var text string

	for scanner.Scan() {
		text += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return text, nil
}
