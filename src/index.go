package src

import (
	"context"
	"log"
	"os"
	"time"

	_ "github.com/daily-utils/iLLM-backend/docs"
	"github.com/daily-utils/iLLM-backend/src/controllers"
	"github.com/daily-utils/iLLM-backend/src/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title iLLM Backend API
// @version 1.0
// @description This is a sample server for iLLM backend.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func Run(ctx context.Context) {
	mongoURL := os.Getenv("MONGO_URL")
	clientOptions := options.Client().ApplyURI(mongoURL)

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

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

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is default base route for go server",
		})
	})

	route.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running!",
		})
	})

	// add routers

	// ask routes
	route.POST("/ask", controllers.Ask)

	// context routes
	route.POST("/context/docx", controllers.ProvideContextForDocx)
	route.POST("/context/link", controllers.ProvideContextForLink)
	route.POST("/context/plaintext", controllers.ProvideContextForPlainText)
	route.POST("/context/txtfile", controllers.ProvideContextForText)
	route.POST("/context/csv", controllers.ProvideContextForCSV)
	route.POST("/context/pdf", controllers.ProvideContextForPdf)

	route.Run(":8090")
}
