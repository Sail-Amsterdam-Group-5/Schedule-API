package main

import (
	"log"
	"os"
	"schedule-api/controller"
	"schedule-api/docs"
	middleware "schedule-api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	router := gin.Default()

	// Initialize Swagger doc info
	docs.SwaggerInfo.BasePath = "/"

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Schedule CRUD routes
	schedule := router.Group("/schedule")
	{
		// object voor location in de task als response
		schedule.POST("/", middleware.CheckScope("admin"), controller.CreateDummyData)                  // only for testing
		schedule.GET("/", middleware.CheckScope("volunteer"), controller.GetAllTasks)                   // Volunteer
		schedule.GET("/:date", middleware.CheckScope("volunteer"), controller.GetSchedule)              // Volunteer
		schedule.GET("/group/:groupid", middleware.CheckScope("volunteer"), controller.GetTasksByGroup) // Volunteer

		schedule.POST("/task", middleware.CheckScope("team-lead"), controller.CreateTask) // Team Lead

		schedule.GET("/task/:id", middleware.CheckScope("volunteer"), controller.GetTask)       // Volunteer
		schedule.PUT("/task/:id", middleware.CheckScope("team-lead"), controller.UpdateTask)    // Team Lead
		schedule.DELETE("/task/:id", middleware.CheckScope("team-lead"), controller.DeleteTask) // Team Lead

		schedule.POST("/task/:id/checkin", middleware.CheckScope("volunteer"), controller.CheckIn)   // Volunteer
		schedule.POST("/task/:id/cancel", middleware.CheckScope("volunteer"), controller.CancelTask) // Volunteer

		schedule.GET("/task/checkins", middleware.CheckScope("team-lead"), controller.GetAllCheckIns)                    // Team Lead
		schedule.GET("/task/checkins/:taskId/:UserId", middleware.CheckScope("volunteer"), controller.GetCheckInForTask) // Volunteer
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
