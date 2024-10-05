package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	// "github.com/daily-utils/iLLM-backend/src/models"
	"github.com/daily-utils/iLLM-backend/src/models"
	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/gin-gonic/gin"
)

type ContextRequestBody struct {
	ContextProvided  string `json:"contextProvider"`
	ContextExtension string `json:"contextExtension"`
	IsLink           bool   `json:"isLink"`
	DomainOfProvider string `json:"domainOfProvider"`
	Model string `json:"model"`
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
	var prompt string
	var body string

	if requestBody.ContextExtension == "pdf" {
		fileContext, err := utils.ExtractTextFromPDF(requestBody.ContextProvided)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		prompt = "I, personally, think that the following text is very interesting. It is like array contains text. Every element in array is text for that page respectively. Please keep it in your context will ask you questions on this context: " + fileContext

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
	} else if requestBody.ContextExtension == "txt" {
		fileContext, err := utils.GetTextFromFile(requestBody.ContextProvided)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		prompt = "I, personally, think that the following text is very interesting. It is text extracted from a txt. Please keep it in your context will ask you questions on this context: " + fileContext

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
	} else if requestBody.ContextExtension == "text" {
		fileContext = requestBody.ContextProvided
		prompt = "This is the text I want you to remember:" + fileContext

		println("prompt: ", prompt)

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
	} else if requestBody.ContextExtension == "docx" {

		fileContext, err := utils.ExtractTextFromDocx(requestBody.ContextProvided)
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
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid context extension"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": string(body)})
}
