package controller

import (
	"net/http"

	"schedule-api/model"
	"schedule-api/service"

	"github.com/gin-gonic/gin"
)

// GetSchedule retreves the schedule for a specific date.
// @Summary Get schedule by date
// @Description Get a schedule by date
// @Param date path string true "Date"
// @Success 200 {object} model.TaskDTO[]
// @Router /schedule/{date} [get]
func GetSchedule(c *gin.Context) {
	date := c.Param("date")
	id := c.Request.BasicAuth("id") // TODO: need to get groupID exually
	// Get the schedule
	schedule, err := service.GetAllTaskForDate(date, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// // Return the schedule
	c.JSON(http.StatusOK, schedule)
}

// GetTasks retreves the tasks for a specific date.
// @Summary Get the tasks by date and group
// @Description Get a list of tasks by date and group
// @Param date path string true "Date"
// @Param groupid path string true "Group ID"
// @Success 200 {object} model.TaskDTO[]
// @Router /schedule/{date}/{groupid} [get]
func GetTasks(c *gin.Context) {
	date := c.Param("date")
	id := c.Request.BasicAuth("id") // TODO: need to get groupID exually
	// Get the schedule
	schedule, err := service.GetAllTaskForDate(date, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// // Return the schedule
	c.JSON(http.StatusOK, schedule)
}

// GetTask retreves a specific Task.
// @Summary Get a task by ID
// @Description Get a task by ID
// @Param id path string true "ID"
// @Success 200 {object} model.TaskDTO
// @Router /schedule/task/{id} [get]
func GetTask(c *gin.Context) {
	taskid := c.Param("ID")

	task := service.GetTaskById(taskid)

	// // Return the schedule
	c.JSON(http.StatusOK, task)
}

// CreateTask creates a new Task.
// @Summary Create a new task
// @Description Create a new task
// @Param task body model.Task true "Task"
// @Success 200 {object} model.TaskDTO
// @Router /schedule/task [post]
func CreateTask(c *gin.Context) {
	schedule := service.CreateTask(c.Request.Body.Read()) // TODO: need more research on this

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// // Return the schedule
	c.JSON(http.StatusOK, schedule)
}

// UpdateTask creates a Task.
// @Summary Updates a task
// @Description Update a task
// @Param id path string true "ID"
// @Param task body model.Task true "Task"
// @Success 200 {object} model.TaskDTO
// @Router /schedule/task/{id} [put]
func UpdateTask(c *gin.Context) {

	id := c.Param("id")
	toUpdate := c.Request.Body.Read() // TODO: need more research on this
	update := service.UpdateTask(toUpdate)

	c.JSON(http.StatusOK, update)
}

// DeleteTask deletes a Task.
// @Summary Deletes a task
// @Description Delete a task
// @Param id path string true "ID"
// @Success 200 string
// @Router /schedule/task/{id} [delete]
func DeleteTask(c *gin.Context) {
	id := c.Param("ID")
	task := service.GetTaskById(id)
	delete := service.DeleteTask(task.PrimaryKey, task.RowKey)

	if delete == false {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
	}
	// Here you would delete the task from the database using the id.
	c.JSON(http.StatusOK, gin.H{"message": "Task with ID " + id + " deleted successfully"})
}

// CheckIn checks in on a Task.
// @Summary CheckIn on a task
// @Description CheckIn on a task
// @Param id path string true "ID"
// @Success 200 {object} model.CheckInDTO
// @Router /schedule/task/{id} [post]
func CheckIn(c *gin.Context) {
	// taskId := c.Param("id")
	// userId :=

	var checkin model.CheckInDTO
	// Here you would upload the checkin to the database.
	c.JSON(http.StatusOK, checkin)
}

// Cancel a Task.
// @Summary Cancel a task
// @Description Cancel a task
// @Param id path string true "ID"
// @Success 200 {object} model.CheckInDTO
// @Router /schedule/task/{id} [patch]
func CancelTask(c *gin.Context) {
	// taskId := c.Param("id")
	// userId :=

	var checkin model.CheckInDTO
	// Here you would upload the cancel to the database.
	c.JSON(http.StatusOK, checkin)
}
