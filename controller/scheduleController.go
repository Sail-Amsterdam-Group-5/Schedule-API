package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"schedule-api/model"
	"schedule-api/service"

	"github.com/gin-gonic/gin"
)

// GetSchedule gets all tasks.
// @Summary Gets all tasks
// @Description Get all Tasks
// @Success 200 {object} model.TaskDTO[]
// @Router /schedule [get]
func GetAllTasks(c *gin.Context) {
	// Get the schedule
	schedule, err := service.GetAllTasks(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the schedule
	c.JSON(http.StatusOK, schedule)
}

// GetSchedule retreves the schedule for a specific date.
// @Summary Get schedule by date
// @Description Get a schedule by date
// @Param date path string true "Date"
// @Success 200 {object} model.TaskDTO[]
// @Router /schedule/{date} [get]
func GetSchedule(c *gin.Context) {
	date := c.Param("date")
	groupId := c.Request.Header.Get("X-User-GroupId")
	// Get the schedule
	schedule, err := service.GetAllTaskForDate(c.Request.Context(), date, groupId)

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
// func GetTasks(c *gin.Context) { // do i use this?
// 	date := c.Param("date")
// 	id := c.Param("groupid")
// 	// Get the schedule
// 	schedule, err := service.GetAllTaskForDate(c.Request.Context(), date, id)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// // Return the schedule
// 	c.JSON(http.StatusOK, schedule)
// }

// GetTasks retreves the tasks for a specific group.
// @Summary Get the tasks for a group
// @Description Get a list of tasks by  group
// @Param groupid path string true "Group ID"
// @Success 200 {object} model.TaskDTO[]
// @Router /schedule/group/{groupid} [get]
func GetTasksByGroup(c *gin.Context) {
	groupid := c.Param("groupid")
	// Get the schedule
	schedule, err := service.GetAllTaskForGroup(c.Request.Context(), groupid)

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
	taskid := c.Param("id")

	task, err := service.GetTaskById(c.Request.Context(), taskid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// // Return the schedule
	c.JSON(http.StatusOK, task)
}

// CreateTask creates a new Task.
// @Summary Create a new task
// @Description Create a new task
// @Param task body model.Task true "Task"
// @Success 200 {object} model.TaskDTO
// @Failure 500 {object} string
// @Router /schedule/task [post]
func CreateTask(c *gin.Context) {
	var task model.Task
	// Bind JSON too the model
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := service.CreateTask(c.Request.Context(), task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// // Return the schedule
	c.JSON(http.StatusOK, gin.H{
		"Message": "Task created succesfully",
		"task":    response,
	})
}

// UpdateTask creates a Task.
// @Summary Updates a task
// @Description Update a task
// @Param id path string true "ID"
// @Param task body model.Task true "Task"
// @Success 200 {object} model.TaskDTO
// @Router /schedule/task/{id} [put]
func UpdateTask(c *gin.Context) {
	var task model.TaskDTO
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update, err := service.UpdateTask(c.Request.Context(), task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, update)
}

// DeleteTask deletes a Task.
// @Summary Deletes a task
// @Description Delete a task
// @Param id path string true "id"
// @Success 200 {object} string
// @Router /schedule/task/{id} [delete]
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task, err := service.GetTaskById(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	delete := service.DeleteTask(c.Request.Context(), task.PrimaryKey, task.RowKey)

	if !delete {
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
// @Router /schedule/task/{id}/checkin [post]
func CheckIn(c *gin.Context) {
	taskId := c.Param("id")
	userId := c.Request.Header.Get("X-User-Id")

	checkin, err := service.Checkin(c.Request.Context(), userId, taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Here you would upload the checkin to the database.
	c.JSON(http.StatusOK, checkin)
}

// Cancel a Task.
// @Summary Cancel a task
// @Description Cancel a task
// @Param id path string true "ID"
// @Success 200 {object} model.CheckInDTO
// @Router /schedule/task/{id}/cancel [post]
func CancelTask(c *gin.Context) {
	taskId := c.Param("id")
	userId := c.Request.Header.Get("X-User-Id")
	var requestBody struct {
		Reason string `json:"reason"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reason := requestBody.Reason
	checkin, err := service.CancelTask(c.Request.Context(), userId, taskId, reason)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Here you would upload the cancel to the database.
	c.JSON(http.StatusOK, checkin)
}

// GetAllCheckIns gets all checkins.
// @Summary Get all checkins
// @Description Get all checkins
// @Success 200 {object} model.CheckInDTO[]
// @Router /schedule/task/checkins [get]
func GetAllCheckIns(c *gin.Context) {
	checkins, err := service.GetAllCheckIns(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, checkins)
}

// GetCheckInForTask gets a checkin for a task.
// @Summary Get a checkin for a task
// @Description Get a checkin for a task
// @Param taskId path string true "Task ID"
// @Param UserId path string true "User ID"
// @Success 200 {boolean} boolean
// @Router /schedule/task/checkins/{taskId}/{UserId} [get]
func GetCheckInForTask(c *gin.Context) {
	taskId := c.Param("taskId")
	userId := c.Param("UserId")

	checkin, err := service.GetCheckInForTask(c.Request.Context(), userId, taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, checkin)
}

// @Summary Create dummy data
// @Description Create dummy data
// @Success 200 {object} string
// @Router /schedule [post]
func CreateDummyData(c *gin.Context) {

	file, err := os.Open("Mockdata.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var tasks []model.Task
	if err := json.Unmarshal(byteValue, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, task := range tasks {
		if _, err := service.CreateTask(c.Request.Context(), task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tasks created successfully"})
}
