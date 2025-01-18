package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"schedule-api/model"
	"schedule-api/repository"

	"github.com/google/uuid"
)

// get all for user
func GetAllTaskForUser(ctx context.Context, userId string) ([]byte, error) {

	tasks := repository.GetAllTaskForUser(ctx, userId)
	return json.Marshal(tasks)
}

// get all for group
func GetAllTaskForGroup(ctx context.Context, groupId string) ([]model.TaskDTO, error) {
	tasks, err := repository.GetAllTaskForGroup(ctx, groupId)
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	}
	return tasks, nil
}

// get all for date
func GetAllTaskForDate(ctx context.Context, date string, groupId string) ([]model.TaskDTO, error) {
	tasks, err := repository.GetAllTaskForDate(ctx, date, groupId)
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	}
	return tasks, nil
}

// get by id
func GetTaskById(ctx context.Context, id string) (model.TaskDTO, error) {
	task, err := repository.GetTaskById(ctx, id)

	if err != nil {
		return model.TaskDTO{}, err
	}
	return task, nil
}

// update a task
func UpdateTask(c context.Context, task model.TaskDTO) (model.Task, error) {
	if repository.UpdateTask(c, task) {

		DbTask, err := repository.GetTaskById(c, task.Id)

		Location, err := GetLocation(DbTask.Utillity)

		taskModel := model.Task{
			Id:          DbTask.Id,
			GroupId:     DbTask.GroupId,
			Name:        DbTask.Name,
			Description: DbTask.Description,
			Date:        DbTask.Date,
			StartTime:   DbTask.StartTime,
			EndTime:     DbTask.EndTime,
			Utillity:    Location,
		}
		if err != nil {
			return model.Task{}, err
		}
		return taskModel, nil
	}
	return model.Task{}, errors.New("failed to update task")
}

// delete a task
func DeleteTask(ctx context.Context, pk string, rk string) bool {
	return repository.DeleteTask(ctx, pk, rk)
}

// create a task
func CreateTask(ctx context.Context, task model.Task) (model.TaskDTO, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	taskDTO := model.TaskDTO{
		PrimaryKey:  task.Date.String() + task.GroupId,
		RowKey:      id.String(),
		Id:          id.String(),
		GroupId:     task.GroupId,
		Name:        task.Name,
		Description: task.Description,
		Date:        task.Date,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
		Utillity:    task.Utillity.Id,
	}

	response, err := repository.CreateTask(ctx, taskDTO)

	return response, err
}

func GetLocation(locationId string) (model.Utillity, error) {
	url := fmt.Sprintf("%s/%s", os.Getenv("MAP_API_URL"), locationId)
	resp, err := http.Get(url)
	if err != nil {
		return model.Utillity{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Utillity{}, errors.New("failed to get location")
	}

	var location model.Utillity
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return model.Utillity{}, err
	}
	return location, nil
}

func GetAllTasks(ctx context.Context) ([]model.TaskDTO, error) {
	tasks, err := repository.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	}
	return tasks, nil
}
