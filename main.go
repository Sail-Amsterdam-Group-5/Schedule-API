package main

import (
	"log"
	"os"
	"schedule-api/controller"
	"schedule-api/docs"
	middleware "schedule-api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
		schedule.GET("/:date", middleware.CheckScope("volunteer"), controller.GetSchedule) // works but date does not work

		schedule.POST("/task", middleware.CheckScope("team-lead"), controller.CreateTask) //

		schedule.GET("/task/:id", middleware.CheckScope("volunteer"), controller.GetTask)       // returns next one in line if id is not found
		schedule.PUT("/task/:id", middleware.CheckScope("team-lead"), controller.UpdateTask)    //
		schedule.DELETE("/task/:id", middleware.CheckScope("team-lead"), controller.DeleteTask) //

		schedule.POST("/task/:id/checkin", middleware.CheckScope("volunteer"), controller.CheckIn)   //
		schedule.POST("/task/:id/cancel", middleware.CheckScope("volunteer"), controller.CancelTask) //

		//		schedule.GET("/task/:id/checkin", middleware.CheckScope("team-lead"), controller.GetCheckInForTask)
		//		schedule.GET("/task/checkins", middleware.CheckScope("team-lead"), controller.GetAllCheckIns)

	}
	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
