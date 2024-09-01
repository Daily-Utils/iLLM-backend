package utils

import (
	"os"

	"code.sajari.com/docconv"
)

func ExtractTextFromDocx(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := f
	r = f

	tmpl, _, err := docconv.ConvertDocx(r)
	if err != nil {
		return "", err
	}

	return tmpl, nil
}
