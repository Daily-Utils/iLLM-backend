package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/daily-utils/iLLM-backend/src/models"
	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/gin-gonic/gin"
)

type ContextRequestBodyForText struct {
	ContextProvided string `json:"contextProvider"`
	Model           string `json:"model"`
}

type ContextRequestBodyForTextFile struct {
	Response string `json:"response"`
}

type ContextRequestBodyForTextFileErr struct {
	Error string `json:"error"`
}

// ProvideContextForText godoc
// @Summary Provide context for text
// @Description Provide context for the model for text
// @Tags context
// @Accept json
// @Produce json
// @Param body body ContextRequestBodyForText true "Request body"
// @Success 200 {object} ContextRequestBodyForTextFile
// @Failure 500 {object} ContextRequestBodyForTextFileErr
// @Router /context/txtfile [post]
func ProvideContextForText(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var requestBody ContextRequestBodyForText
	if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fileContext, err := utils.GetTextFromFile(requestBody.ContextProvided)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	prompt := "I, personally, think that the following text is very interesting. It is text extracted from a txt. Please keep it in your context will ask you questions on this context: " + fileContext

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
