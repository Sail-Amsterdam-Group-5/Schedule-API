package main

import (
	middleware "schedule-api/Middleware"
	"schedule-api/controller"
	"schedule-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()

	// Initialize Swagger doc info
	docs.SwaggerInfo.BasePath = "/"

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Schedule CRUD routes
	schedule := router.Group("/schedule")
	{
		schedule.GET("/:date", middleware.CheckScope("volunteer"), controller.GetSchedule)
		schedule.GET("/:date/:groupid", middleware.CheckScope("team-lead"), controller.GetSchedule)

		schedule.POST("/task", middleware.CheckScope("team-lead"), controller.CreateTask)

		schedule.GET("/task/:id", middleware.CheckScope("volunteer"), controller.GetTask)
		schedule.PUT("/task/:id", middleware.CheckScope("team-lead"), controller.UpdateTask)
		schedule.DELETE("/task/:id", middleware.CheckScope("team-lead"), controller.DeleteTask)
		schedule.POST("/task/:id", middleware.CheckScope("volunteer"), controller.CheckIn)
		schedule.PATCH("/task/:id", middleware.CheckScope("volunteer"), controller.CancelTask)
	}

	router.Run(":8080")
}
