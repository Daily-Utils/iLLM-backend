package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/daily-utils/iLLM-backend/src/models"
	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/gin-gonic/gin"
)

type AskRequestBody struct {
	Prompt  string  `json:"prompt"`
	Model   string  `json:"model"`
	Context []int64 `json:"context"`
}

// Ask godoc
// @Summary Ask a question
// @Description Ask a question and get a response
// @Tags ask
// @Accept json
// @Produce json
// @Param body body AskRequestBody true "Request body"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.ResponseError
// @Router /temp/ask [post]
func (ctrl *Controller) TempAsk(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var requestBody AskRequestBody
	if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Extract prompt from body
	prompt := models.Ask{
		Model:   requestBody.Model,
		Prompt:  string(requestBody.Prompt),
		Stream:  false,
		Context: requestBody.Context,
	}

	body, err := utils.RequestClient(prompt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response models.Response
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := utils.MongoInsertResponse(c, ctrl.MongoClient, response, "iLLM", "temp_responses"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}
