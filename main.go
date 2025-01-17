package main

import (
	"log"
	"os"
	"schedule-api/controller"
	"schedule-api/docs"
	middleware "schedule-api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error loading .env file: %v", err)
	//}
	// Check for OC secrets
	log.Println("Env test: %v", os.Getenv("HTTP_PLATFORM_PORT"))

	router := gin.Default()

	// Initialize Swagger doc info
	docs.SwaggerInfo.BasePath = "/"

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Schedule CRUD routes
	schedule := router.Group("/schedule")
	{
		// object voor location in de task als response
		schedule.POST("/", middleware.CheckScope("admin"), controller.CreateDummyData)     // only for testing
		schedule.GET("/", middleware.CheckScope("volunteer"), controller.GetAllTasks)      //
		schedule.GET("/:date", middleware.CheckScope("volunteer"), controller.GetSchedule) //

		schedule.POST("/task", middleware.CheckScope("team-lead"), controller.CreateTask) //

		schedule.GET("/task/:id", middleware.CheckScope("volunteer"), controller.GetTask)       //
		schedule.PUT("/task/:id", middleware.CheckScope("team-lead"), controller.UpdateTask)    //
		schedule.DELETE("/task/:id", middleware.CheckScope("team-lead"), controller.DeleteTask) //

		schedule.POST("/task/:id/checkin", middleware.CheckScope("volunteer"), controller.CheckIn)   //
		schedule.POST("/task/:id/cancel", middleware.CheckScope("volunteer"), controller.CancelTask) //

		//		schedule.GET("/task/:id/checkin", middleware.CheckScope("team-lead"), controller.GetCheckInForTask)
		//		schedule.GET("/task/checkins", middleware.CheckScope("team-lead"), controller.GetAllCheckIns)

	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
