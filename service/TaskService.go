package service

import (
	"encoding/json"
	"schedule-api/model"
	"schedule-api/repository"
)

// get all for user
func GetAllTaskForUser(userId string) ([]byte, error) {

	tasks := repository.GetAllTaskForUser(userId)
	return json.Marshal(tasks)
}

// get all for date
func GetAllTaskForDate(date string) {
	return repository.GetAllTaskForDate(date)
}

// get by id
func GetTaskById(id string) {
	return repository.GetTaskById(id)
}

// update a task
func UpdateTask(task model.TaskDTO) {
	id := task.Id

	return repository.UpdateTask(id, task)
}

// delete a task
func DeleteTask(id string) {
	return repository.DeleteTask(id)
}

// create a task
func CreateTask(task model.Task) {
	return repository.CreateTask(task)
}
