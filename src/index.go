package src

import (
	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/gin-gonic/gin"
)

func Run() {
	utils.LoadEnv()

	route := gin.Default()
	route.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	route.Use(utils.Logger())

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is default base route for go server",
		})
	})

	route.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running",
		})
	})

	// add routers

	route.Run(":8090")
}
