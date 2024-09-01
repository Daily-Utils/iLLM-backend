package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/daily-utils/iLLM-backend/src/models"
	"github.com/gin-gonic/gin"
)

type AskRequestBody struct {
	Prompt string `json:"prompt"`
}

type AskResponseBody struct {
	Response string `json:"response"`
}

type AskErrorResponseBody struct {
	Error string `json:"error"`
}

// Ask godoc
// @Summary Ask a question
// @Description Ask a question and get a response
// @Tags ask
// @Accept json
// @Produce json
// @Param body body AskRequestBody true "Request body"
// @Success 200 {object} AskResponseBody
// @Failure 500 {object} AskErrorResponseBody
// @Router /ask [post]
func Ask(c *gin.Context) {
	// extract variables from response body
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var requestBody AskRequestBody
	if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Extract prompt from body
	prompt := models.Ask{
		Model:  "llama3",
		Prompt: string(requestBody.Prompt),
		Stream: false,
	}

	promptJSON, err := json.Marshal(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{
		Timeout: time.Minute * 5,
	}

	LLAMA_API := os.Getenv("LLAMA_API")

	// hit the api
	req, err := http.NewRequest("POST", LLAMA_API+"api/generate", bytes.NewBuffer(promptJSON))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate response"})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Use the prompt variable as needed
	c.JSON(http.StatusOK, gin.H{"response": string(body)})
}
