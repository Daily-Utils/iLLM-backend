package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	// "github.com/daily-utils/iLLM-backend/src/models"
	"github.com/gin-gonic/gin"
)

type ContextRequestBody struct {
	FilePath string `json:"filePath"`
}

func ProvideContext(c *gin.Context){
	bodyBytes , err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var requestBody ContextRequestBody
	if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// hit the api

	c.JSON(http.StatusOK, gin.H{"filePath": requestBody.FilePath})
}