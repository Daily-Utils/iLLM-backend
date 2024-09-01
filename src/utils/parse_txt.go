package utils

import (
	"bufio"
	"os"
)

func GetTextFromFile(filePath string) (string, error) {
	file, err := os.Open((filePath))
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
