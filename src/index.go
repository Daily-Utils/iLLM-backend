package src

import (
	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/gin-gonic/gin"
)

func Run() {
	utils.LoadEnv()

	route := gin.Default()
	route.Use(utils.Logger())

	route.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running",
		})
	})

	// add routers

	route.Run(":8090")
}
