package utils

import (
	"bytes"
	"fmt"
    "github.com/pdfcpu/pdfcpu/pkg/api"
    "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"os"
)

func ParsePDFToString(filePath string) (string, error) {
    // Read the PDF file
    fileBytes, err := os.ReadFile(filePath)
    if err != nil {
        return "", fmt.Errorf("failed to read file: %v", err)
    }

    // Create a PDF context
    ctx, err := api.ReadContext(bytes.NewReader(fileBytes), pdfcpu.DefaultImportConfig())
    if err != nil {
        return "", fmt.Errorf("failed to create PDF context: %v", err)
    }

    // Extract content from the PDF
    var buf bytes.Buffer
    if err := api.ExtractContent(ctx, &buf); err != nil {
        return "", fmt.Errorf("failed to extract content: %v", err)
    }

    return buf.String(), nil
}
