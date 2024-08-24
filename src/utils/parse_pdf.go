package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func ParsePDFToString(filePath string) (string, error) {
	// Read the PDF file
    file, err := os.Open(filePath)
    if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
    }
	defer file.Close()

	pdfReader, err := model.NewPdfReader(file)
	if err != nil {
		return "", fmt.Errorf("failed to create pdf reader: %v", err)
	}

	numPages, err := pdfReader.GetNumPages()

	if err != nil {
		return "", fmt.Errorf("failed to get number of pages: %v", err)
	}

	text := ""
	textArr := []string{}
	for i := 0; i < min(numPages, 60); i++ {
		pageNum := i + 1
		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return "", fmt.Errorf("failed to get page: %v", err)
		}
		ex, err := extractor.New(page)
		if err != nil {
			return "", fmt.Errorf("failed to create text extractor: %v", err)
		}
		pageText, err := ex.ExtractText()
		if err != nil {
			return "", fmt.Errorf("failed to extract text: %v", err)
		}
		text += pageText
		textArr = append(textArr, pageText)
		text = ""
	}

	jsonText, err := json.Marshal(textArr)
	if err != nil {
		return "", fmt.Errorf("failed to marshal text: %v", err)
	}
	
	return string(jsonText), nil
}
