package utils

import (
    "bytes"
    "fmt"
    "github.com/martoche/pdf"
)

func ExtractTextFromPDF(filePath string) (string, error) {
    // Open the PDF file
    r, err := pdf.Open(filePath)
    if err != nil {
        return "", fmt.Errorf("failed to open PDF file: %v", err)
    }

    // Get plain text from the PDF
    p, err := r.GetPlainText()
    if err != nil {
        return "", fmt.Errorf("failed to get plain text from PDF: %v", err)
    }

    // Check the type assertion
    buf, ok := p.(*bytes.Buffer)
    if !ok {
        return "", fmt.Errorf("the library no longer uses bytes.Buffer to implement io.Reader")
    }

    // Return the extracted text as a string
    return buf.String(), nil
}

