package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/daily-utils/iLLM-backend/src/models"
	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/gin-gonic/gin"
)

// ProvideContextForPdf godoc
// @Summary Provide context for PDF
// @Description Provide context for the model for PDF
// @Tags context
// @Accept json
// @Produce json
// @Param body body RequestBodyForPDF true "Request body"
// @Success 200 {object} ResponseBodyForPDF
// @Failure 500 {object} ErrorResponseBodyForPDF
// @Router /context/pdf [post]
func ProvideContextForPdf(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var requestBody models.RequestBodyForPDF
	if err := json.Unmarshal(bodyBytes, &requestBody); 
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	prompt := "I, personally, think that the following text is very interesting. Please keep it in your context will ask you questions on this context: " + requestBody.Body.Data

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