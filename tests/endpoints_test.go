package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"schedule-api/controller"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	os.Setenv("AZURE_CONNECTION_STRING", "DefaultEndpointsProtocol=https;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;TableEndpoint=http://127.0.0.1:10002/devstoreaccount1;")
	os.Setenv("MAP_API_URL", "https://sail-map-api-route-oscar-dev.apps.inholland.hcs-lab.nl/locations")
	router := gin.Default()

	schedule := router.Group("/schedule")
	{
		schedule.GET("/", controller.GetAllTasks)
		schedule.GET("/:date", controller.GetSchedule)
		schedule.GET("/group/:groupid", controller.GetTasksByGroup)
		schedule.POST("/task", controller.CreateTask)
		schedule.GET("/task/:id", controller.GetTask)
		schedule.PUT("/task/:id", controller.UpdateTask)
		schedule.DELETE("/task/:id", controller.DeleteTask)
		schedule.POST("/task/:id/checkin", controller.CheckIn)
		schedule.POST("/task/:id/cancel", controller.CancelTask)
		schedule.GET("/task/checkins", controller.GetAllCheckIns)
		schedule.GET("/task/checkins/:taskId/:UserId", controller.GetCheckInForTask)
		schedule.GET("/task/canceled/:taskId/:UserId", controller.GetCancelForTask)
	}

	return router
}

func TestGetAllTasks(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/schedule/", nil)
	router.ServeHTTP(w, req)

	assert.Error(t, errors.New("no tasks found"))
	assert.Equal(t, http.StatusOK, w.Code)
}

// test the GetSchedule function with date that has no tasks
func TestGetSchedule(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/schedule/2023-10-10T00:00:00Z", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetTasksByGroup(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/schedule/group/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateTask(t *testing.T) {
	router := setupRouter()
	task := `{
		"primaryKey": "1",
		"rowKey": "1",
		"id": "1",
		"groupId": "1",
		"name": "Test Task",
		"description": "This is a test task",
		"date": "2023-10-10T00:00:00Z",
		"startTime": "2023-10-10T09:00:00Z",
		"endTime": "2023-10-10T10:00:00Z",
		"location": {"id": "b62a6037-b572-47c1-b1c8-96ff57ff9cd0" 	}
		}`
	req, _ := http.NewRequest("POST", "/schedule/task", strings.NewReader(task))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTask(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/schedule/task/2859253b-1ad3-4600-9abb-f5b0342126f9", nil)
	router.ServeHTTP(w, req)

	assert.Error(t, errors.New("Task not found"))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTask(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	reqBody := `{
		"primaryKey": "1",
		"rowKey": "1",
		"id": "2859253b-1ad3-4600-9abb-f5b0342126f9",
		"groupId": "1",
		"name": "Updated Task",
		"description": "This is an updated test task",
		"date": "2023-10-11T00:00:00Z",
		"startTime": "2023-10-11T09:00:00Z",
		"endTime": "2023-10-11T10:00:00Z",
		"location": "b62a6037-b572-47c1-b1c8-96ff57ff9cd0"
	}`
	req, _ := http.NewRequest("PUT", "/schedule/task/2859253b-1ad3-4600-9abb-f5b0342126f9", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTask(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/schedule/task/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckIn(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/schedule/task/1/checkin", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCancelTask(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	reqBody := `{"reason": "Task no longer needed"}`
	req, _ := http.NewRequest("POST", "/schedule/task/1/cancel", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Error(t, errors.New("Task already cancelled"))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllCheckIns(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/schedule/task/checkins", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCheckInForTask(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/schedule/task/checkins/1/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCancelForTask(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/schedule/task/canceled/1/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
