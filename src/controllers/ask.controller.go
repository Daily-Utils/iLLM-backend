package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	// "github.com/daily-utils/iLLM-backend/src/models"
	"github.com/daily-utils/iLLM-backend/src/models"
	"github.com/gin-gonic/gin"
)

type AskRequestBody struct {
	Prompt string `json:"prompt"`
}

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

	// Use the prompt variable as needed
	c.JSON(http.StatusOK, gin.H{"prompt": prompt})
}
