package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	// "github.com/daily-utils/iLLM-backend/src/models"
	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/gin-gonic/gin"
)

type ContextRequestBody struct {
	ContextProvided  string `json:"contextProvider"`
	ContextExtension string `json:"contextExtension"`
	IsLink           bool   `json:"isLink"`
	DomainOfProvider string `json:"domainOfProvider"`
}

type ContextResponseBody struct {
	Response string `json:"response"`
}

type ContextErrorResponseBody struct {
	Error string `json:"error"`
}

// Context godoc
// @Summary Provide context
// @Description Provide context for the model
// @Tags context
// @Accept json
// @Produce json
// @Param body body ContextRequestBody true "Request body"
// @Success 200 {object} ContextResponseBody
// @Failure 500 {object} ContextErrorResponseBody
// @Router /context [post]
func ProvideContext(c *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var requestBody ContextRequestBody
	if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// check for kind of context
	if requestBody.IsLink {
		linkContext, err := utils.GetTextFromLink(requestBody.ContextProvided, requestBody.DomainOfProvider)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": linkContext})
	}

	var fileContext string

	if requestBody.ContextExtension == "pdf" {
		fileContext, err := utils.ExtractTextFromPDF(requestBody.ContextProvided)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": fileContext})
	} else if requestBody.ContextExtension == "txt" {
		fileContext, err := utils.GetTextFromFile(requestBody.ContextProvided)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": fileContext})
	} else if requestBody.ContextExtension == "text" {
		c.JSON(http.StatusOK, gin.H{"response": requestBody.ContextProvided})
	} else if requestBody.ContextExtension == "docx" {
		fileContext, err := utils.ExtractTextFromDocx(requestBody.ContextProvided)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": fileContext})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid context extension"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": fileContext})
}
