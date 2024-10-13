package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/daily-utils/iLLM-backend/src/models"
	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/gin-gonic/gin"
)

// TODO: Impliment array for DomainOfProvider
type ContextRequestBodyForLink struct {
	Link             string `json:"Link"`
	ContextExtension string `json:"contextExtension"`
	DomainOfProvider string `json:"domainOfProvider"`
	Model            string `json:"model"`
}

type ContextResponseBodyForLink struct {
	Response string `json:"response"`
}

type ContextErrorResponseBodyForLink struct {
	Error string `json:"error"`
}

// ProvideContextForLink godoc
// @Summary Provide context for link
// @Description Provide context for the model for link
// @Tags context
// @Accept json
// @Produce json
// @Param body body ContextRequestBodyForLink true "Request body"
// @Success 200 {object} ContextResponseBodyForLink
// @Failure 500 {object} ContextErrorResponseBodyForLink
// @Router /context/link [post]
func ProvideContextForLink(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var requestBody ContextRequestBodyForLink
	if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var linkContext string
	linkContextText, err := utils.GetTextFromLink(requestBody.Link, requestBody.DomainOfProvider)

	linkContext = linkContextText

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	prompt := "This is web extracted text via a scraper: " + linkContext

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
	if err := json.Unmarshal([]byte(body), &response); 

	err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}
