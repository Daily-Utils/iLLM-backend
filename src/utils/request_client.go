package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/daily-utils/iLLM-backend/src/models"
)

func RequestClient(promptModel models.Ask) (string, error) {
	LLAMA_URL := os.Getenv("LLAMA_URL")

	client := &http.Client{
		Timeout: time.Minute * 10,
	}

	promptJSON, err := json.Marshal(promptModel)

	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	req, err := http.NewRequest("POST", LLAMA_URL+"api/generate", bytes.NewBuffer(promptJSON))
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return string(body), err
}