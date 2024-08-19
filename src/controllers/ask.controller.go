package controllers

import (
	"io/ioutil"
	"net/http"

	// "github.com/daily-utils/iLLM-backend/src/models"
	"github.com/gin-gonic/gin"
)

func Ask(c *gin.Context){
	// extract variables from response body
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	prompt := string(bodyBytes)

	// generate response
	c.JSON(http.StatusOK, gin.H{"prompt": prompt})
}