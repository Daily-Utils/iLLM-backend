package controllers

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/daily-utils/iLLM-backend/src/models"
	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/gin-gonic/gin"
)

// ProvideContextForText godoc
// @Summary Provide context for text
// @Description Provide context for the model for text
// @Tags context
// @Accept json
// @Accept multipart/form-data
// @Produce json
// @Param body formData models.ContextRequestBodyForText true "Request body"
// @Param file formData file true "File to upload"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.ResponseError
// @Router /context/txtfile [post]
func (ctrl *Controller) ProvideContextForText(c *gin.Context) {
	var requestBody models.ContextRequestBodyForText

	if c.ContentType() == "multipart/form-data" {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ext := strings.ToLower(filepath.Ext(file.Filename))

		allowedExtensions := map[string]bool{
			".txt":  true,
			".text": true,
		}

		if !allowedExtensions[ext] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file extension"})
			return
		}

		requestBody.Model = c.PostForm("model")

		fileContext, err := utils.GetTextFromFile(file)
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

		c.JSON(http.StatusOK, gin.H{"response": response})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
	}
}
