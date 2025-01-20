package main

import (
	"log"
	"os"
	"schedule-api/controller"
	"schedule-api/docs"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

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
		schedule.GET("/", controller.GetAllTasks)                   // Volunteer
		schedule.GET("/:date", controller.GetSchedule)              // Volunteer
		schedule.GET("/group/:groupid", controller.GetTasksByGroup) // Volunteer

		schedule.POST("/task", controller.CreateTask) // Team Lead

		schedule.GET("/task/:id", controller.GetTask)       // Volunteer
		schedule.PUT("/task/:id", controller.UpdateTask)    // Team Lead
		schedule.DELETE("/task/:id", controller.DeleteTask) // Team Lead

		schedule.POST("/task/:id/checkin", controller.CheckIn)   // Volunteer
		schedule.POST("/task/:id/cancel", controller.CancelTask) // Volunteer

		schedule.GET("/task/checkins", controller.GetAllCheckIns)                    // Team Lead
		schedule.GET("/task/checkins/:taskId/:UserId", controller.GetCheckInForTask) // Volunteer
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
