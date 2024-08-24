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
	ContextProvided string `json:"contextProvider"`
	IsLink          bool   `json:"isLink"`
	DomainOfProvider string `json:"domainOfProvider"`
}

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

	fileContext, err := utils.ExtractTextFromPDF(requestBody.ContextProvided)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": fileContext})
}
