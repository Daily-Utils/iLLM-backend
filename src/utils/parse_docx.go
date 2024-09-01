package utils

import (
	"fmt"

	"code.sajari.com/docconv"
)

func ExtractTextFromDocx(filePath string) (string, error) {
	doc, err := docconv.ConvertPath(filePath)
	if err != nil {
		return "", err
	}

	var text string
	for _, page := range doc.Body {
		text += fmt.Sprintf("%s\n", string(page))
	}

	return text, nil
}
