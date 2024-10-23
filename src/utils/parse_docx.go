package utils

import (
	"mime/multipart"

	"code.sajari.com/docconv"
)

func ExtractTextFromDocx(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	defer file.Close()

	r := file
	r = file

	tmpl, _, err := docconv.ConvertDocx(r)
	if err != nil {
		return "", err
	}

	return tmpl, nil
}
