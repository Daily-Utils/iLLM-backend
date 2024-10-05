package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/daily-utils/iLLM-backend/src/models"
	"github.com/daily-utils/iLLM-backend/src/utils"
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
		Model:  "llama3.2",
		Prompt: string(requestBody.Prompt),
		Stream: false,
	}

	body, err := utils.RequestClient(prompt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": string(body)})
}
