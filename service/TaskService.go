package service

import (
	"encoding/json"
	"schedule-api/model"
	"schedule-api/repository"
	"strconv"

	"github.com/google/uuid"
)

// get all for user
func GetAllTaskForUser(userId string) ([]byte, error) {

	tasks := repository.GetAllTaskForUser(userId)
	return json.Marshal(tasks)
}

// get all for date
func GetAllTaskForDate(date string, groupId string) ([]byte, error) {
	tasks := repository.GetAllTaskForDate(date, groupId)
	return json.Marshal(tasks)
}

// get by id
func GetTaskById(id string) model.TaskDTO {
	return repository.GetTaskById(id)
}

// update a task
func UpdateTask(task model.TaskDTO) bool {
	return repository.UpdateTask(task)
}

// delete a task
func DeleteTask(pk string, rk string) bool {
	return repository.DeleteTask(pk, rk)
}

// create a task
func CreateTask(task model.Task) model.TaskDTO {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	taskDTO := model.TaskDTO{
		PrimaryKey:  task.Date + strconv.Itoa(task.GroupId),
		RowKey:      task.StartTime + id.String(),
		Id:          id.String(),
		GroupId:     task.GroupId,
		Name:        task.Name,
		Description: task.Description,
		Date:        task.Date,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
		Location:    task.Location,
	}

	repository.CreateTask(taskDTO)

	return taskDTO
}
