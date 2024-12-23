package service

import (
	"context"
	"encoding/json"
	"schedule-api/model"
	"schedule-api/repository"
	"strconv"

	"github.com/google/uuid"
)

// get all for user
func GetAllTaskForUser(ctx context.Context, userId string) ([]byte, error) {

	tasks := repository.GetAllTaskForUser(ctx, userId)
	return json.Marshal(tasks)
}

// get all for date
func GetAllTaskForDate(ctx context.Context, date string, groupId string) ([]byte, error) {
	tasks := repository.GetAllTaskForDate(ctx, date, groupId)
	return json.Marshal(tasks)
}

// get by id
func GetTaskById(ctx context.Context, id string) model.TaskDTO {
	return repository.GetTaskById(ctx, id)
}

// update a task
func UpdateTask(c context.Context, task model.TaskDTO) bool {
	return repository.UpdateTask(c, task)
}

// delete a task
func DeleteTask(ctx context.Context, pk string, rk string) bool {
	return repository.DeleteTask(ctx, pk, rk)
}

// create a task
func CreateTask(ctx context.Context, task model.Task) model.TaskDTO {
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

	repository.CreateTask(ctx, taskDTO)

	return taskDTO
}
