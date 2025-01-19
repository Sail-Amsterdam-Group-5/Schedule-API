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
func GetAllTaskForGroup(ctx context.Context, groupId string) ([]model.Task, error) {
	tasks, err := repository.GetAllTaskForGroup(ctx, groupId)
	if err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	} else {
		var taskModel []model.Task
		for _, task := range tasks {
			Location, err := GetLocation(task.Location)
			if err != nil {
				return nil, err
			}
			taskModel = append(taskModel, model.Task{
				Id:          task.Id,
				GroupId:     task.GroupId,
				Name:        task.Name,
				Description: task.Description,
				Date:        task.Date,
				StartTime:   task.StartTime,
				EndTime:     task.EndTime,
				Location:    Location,
			})
		}
		return taskModel, nil
	}
}

// get all for date
func GetAllTaskForDate(ctx context.Context, date string, groupId string) ([]model.Task, error) {
	tasks, err := repository.GetAllTaskForDate(ctx, date, groupId)
	if err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	} else {
		var taskModel []model.Task
		for _, task := range tasks {
			Location, err := GetLocation(task.Location)
			if err != nil {
				return nil, err
			}
			taskModel = append(taskModel, model.Task{
				Id:          task.Id,
				GroupId:     task.GroupId,
				Name:        task.Name,
				Description: task.Description,
				Date:        task.Date,
				StartTime:   task.StartTime,
				EndTime:     task.EndTime,
				Location:    Location,
			})
		}
		return taskModel, nil
	}
}

// get by id
func GetTaskById(ctx context.Context, id string) (model.Task, error) {
	task, err := repository.GetTaskById(ctx, id)

	Location, err := GetLocation(task.Location)

	taskModel := model.Task{
		Id:          task.Id,
		GroupId:     task.GroupId,
		Name:        task.Name,
		Description: task.Description,
		Date:        task.Date,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
		Location:    Location,
	}
	if err != nil {
		return model.Task{}, err
	}
	return taskModel, nil

}

// update a task
func UpdateTask(c context.Context, task model.TaskDTO) (model.Task, error) {
	if repository.UpdateTask(c, task) {

		DbTask, err := repository.GetTaskById(c, task.Id)

		Location, err := GetLocation(DbTask.Location)

		taskModel := model.Task{
			Id:          DbTask.Id,
			GroupId:     DbTask.GroupId,
			Name:        DbTask.Name,
			Description: DbTask.Description,
			Date:        DbTask.Date,
			StartTime:   DbTask.StartTime,
			EndTime:     DbTask.EndTime,
			Location:    Location,
		}
		if err != nil {
			return model.Task{}, err
		}
		return taskModel, nil
	}
	return model.Task{}, errors.New("failed to update task")
}

// delete a task
func DeleteTask(ctx context.Context, id string) bool {
	return repository.DeleteTask(ctx, id)
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
		Location:    task.Location.Id,
	}

	response, err := repository.CreateTask(ctx, taskDTO)

	return response, err
}

func GetLocation(locationId string) (model.LocationDTO, error) {
	url := fmt.Sprintf("%s/%s", os.Getenv("MAP_API_URL"), locationId)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return model.LocationDTO{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.LocationDTO{}, errors.New("failed to get location")
	}

	var location model.LocationDTO
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return model.LocationDTO{}, err
	}
	return location, nil
}

func GetAllTasks(ctx context.Context) ([]model.Task, error) {
	tasks, err := repository.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	} else {
		var taskModel []model.Task
		for _, task := range tasks {
			Location, err := GetLocation(task.Location)
			if err != nil {
				return nil, err
			}
			taskModel = append(taskModel, model.Task{
				Id:          task.Id,
				GroupId:     task.GroupId,
				Name:        task.Name,
				Description: task.Description,
				Date:        task.Date,
				StartTime:   task.StartTime,
				EndTime:     task.EndTime,
				Location:    Location,
			})
		}
		return taskModel, nil
	}
}
