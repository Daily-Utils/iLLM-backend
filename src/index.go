package src

import (
	"time"

	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/daily-utils/iLLM-backend/src/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	utils.LoadEnv()

	route := gin.Default()
	route.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	route.Use(utils.Logger())

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
	route.POST("/ask", controllers.Ask)

	route.Run(":8090")
}
