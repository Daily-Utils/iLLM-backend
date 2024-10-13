package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/daily-utils/iLLM-backend/src/models"
	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/gin-gonic/gin"
)

type ContextRequestBodyForPlainText struct {
	Text  string `json:"text"`
	Model string `json:"model"`
}

type ContextResponseBodyForPlainText struct {
	Response string `json:"response"`
}

type ContextErrorResponseBodyForPlainText struct {
	Error string `json:"error"`
}

// ProvideContextForPlainText godoc
// @Summary Provide context for plain text
// @Description Provide context for the model for plain text
// @Tags context
// @Accept json
// @Produce json
// @Param body body ContextRequestBodyForPlainText true "Request body"
// @Success 200 {object} ContextResponseBodyForPlainText
// @Failure 500 {object} ContextErrorResponseBodyForPlainText
// @Router /context/plaintext [post]
func ProvideContextForPlainText(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var requestBody ContextRequestBodyForPlainText
	if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	prompt := "I, personally, think that the following text is very interesting. Please keep it in your context will ask you questions on this context: " + requestBody.Text

	promptModel := models.Ask{
			Model:  requestBody.Model,
			Prompt: prompt,
			Stream: false,
		}

		bodyContent, err := utils.RequestClient(promptModel)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}


	body := bodyContent
	var response models.Response

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response.Response})
}