package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/daily-utils/iLLM-backend/src/models"
	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/gin-gonic/gin"
)

type ContextRequestBodyForDocx struct {
	ContextProvided string `json:"contextProvider"`
	Model           string `json:"model"`
}

type ContextResponseBodyForDocx struct {
	Response string `json:"response"`
}

type ContextErrorResponseBodyForDocx struct {
	Error string `json:"error"`
}

// ProvideContextForDocx godoc
// @Summary Provide context for docx
// @Description Provide context for the model for docx
// @Tags context
// @Accept json
// @Produce json
// @Param body body ContextRequestBodyForDocx true "Request body"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.ResponseError
// @Router /context/docx [post]
func ProvideContextForDocx(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var requestBody ContextRequestBodyForDocx
	if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var fileContext string
	var prompt string
	var body string

	fileContext, err = utils.ExtractTextFromDocx(requestBody.ContextProvided)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	prompt = "I, personally, think that the following text is very interesting. Please keep it in your context will ask you questions on this context: " + fileContext

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

	body = bodyContent

	var response models.Response

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}
