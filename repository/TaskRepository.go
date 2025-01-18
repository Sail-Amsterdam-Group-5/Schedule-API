package repository

import (
	"context"
	"errors"
	"log"
	"schedule-api/database"
	"schedule-api/model"
	"time"
)

func parseTime(dateStr string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		log.Printf("Error parsing time: %v", err)
		return time.Time{}
	}
	return parsedTime
}

func GetAllTasks(ctx context.Context) ([]model.TaskDTO, error) {
	entities, err := database.ReadAll(ctx, "Tasks")
	if err != nil {
		return nil, err
	}

	var tasks []model.TaskDTO
	for _, entity := range entities {
		task := model.TaskDTO{
			PrimaryKey:  entity.PartitionKey,
			RowKey:      entity.RowKey,
			Id:          entity.Properties["Id"].(string),
			GroupId:     entity.Properties["GroupId"].(string),
			Name:        entity.Properties["Name"].(string),
			Description: entity.Properties["Description"].(string),
			Date:        parseTime(entity.Properties["Date"].(string)),
			StartTime:   parseTime(entity.Properties["StartTime"].(string)),
			EndTime:     parseTime(entity.Properties["EndTime"].(string)),
			Utillity:    entity.Properties["Location"].(string),
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// get all for user
func GetAllTaskForUser(ctx context.Context, groupId string) []model.TaskDTO {
	filter := database.BuildFilter("GroupId", groupId)

	entities, err := database.ReadFilter(ctx, "Tasks", filter)
	if err != nil {
		return nil
	}

	var tasks []model.TaskDTO
	for _, entity := range entities {
		task := model.TaskDTO{
			Id:          entity.Properties["Id"].(string),
			GroupId:     entity.Properties["GroupId"].(string),
			Name:        entity.Properties["Name"].(string),
			Description: entity.Properties["Description"].(string),
			Date:        parseTime(entity.Properties["Date"].(string)),
			StartTime:   parseTime(entity.Properties["StartTime"].(string)),
			EndTime:     parseTime(entity.Properties["EndTime"].(string)),
			Utillity:    entity.Properties["Location"].(string),
		}
		tasks = append(tasks, task)
	}
	return tasks
}

// get all for group
func GetAllTaskForGroup(ctx context.Context, groupId string) ([]model.TaskDTO, error) {
	filter := database.BuildFilter("GroupId", groupId)
	entities, err := database.ReadFilter(ctx, "Tasks", filter)
	if err != nil {
		return nil, err
	}

	var tasks []model.TaskDTO
	for _, entity := range entities {
		task := model.TaskDTO{
			Id:          entity.Properties["Id"].(string),
			GroupId:     entity.Properties["GroupId"].(string),
			Name:        entity.Properties["Name"].(string),
			Description: entity.Properties["Description"].(string),
			Date:        parseTime(entity.Properties["Date"].(string)),
			StartTime:   parseTime(entity.Properties["StartTime"].(string)),
			EndTime:     parseTime(entity.Properties["EndTime"].(string)),
			Utillity:    entity.Properties["Location"].(string),
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// get all for date
func GetAllTaskForDate(ctx context.Context, date string, groupId string) ([]model.TaskDTO, error) {
	filter := database.BuildDuoFilter("Date", date, "GroupId", groupId)
	entities, err := database.ReadFilter(ctx, "Tasks", filter)
	if err != nil {
		return nil, err
	}

	var tasks []model.TaskDTO
	for _, entity := range entities {
		task := model.TaskDTO{
			PrimaryKey:  entity.PartitionKey,
			RowKey:      entity.RowKey,
			Id:          entity.Properties["Id"].(string),
			GroupId:     entity.Properties["GroupId"].(string),
			Name:        entity.Properties["Name"].(string),
			Description: entity.Properties["Description"].(string),
			Date:        parseTime(entity.Properties["Date"].(string)),
			StartTime:   parseTime(entity.Properties["StartTime"].(string)),
			EndTime:     parseTime(entity.Properties["EndTime"].(string)),
			Utillity:    entity.Properties["Location"].(string),
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// get by id
func GetTaskById(ctx context.Context, id string) (model.TaskDTO, error) {
	filter := database.BuildFilter("Id", id)
	task, err := database.ReadFilter(ctx, "Tasks", filter)
	if err != nil {
		return model.TaskDTO{}, err
	}

	if len(task) > 0 {
		taskDTO := model.TaskDTO{
			PrimaryKey:  task[0].PartitionKey,
			RowKey:      task[0].RowKey,
			Id:          task[0].Properties["Id"].(string),
			GroupId:     task[0].Properties["GroupId"].(string),
			Name:        task[0].Properties["Name"].(string),
			Description: task[0].Properties["Description"].(string),
			Date:        parseTime(task[0].Properties["Date"].(string)),
			StartTime:   parseTime(task[0].Properties["StartTime"].(string)),
			EndTime:     parseTime(task[0].Properties["EndTime"].(string)),
			Utillity:    task[0].Properties["Location"].(string),
		}
		return taskDTO, nil
	}
	return model.TaskDTO{}, errors.New("task not found")
}

// update a task
func UpdateTask(c context.Context, task model.TaskDTO) bool {
	taskMap := map[string]interface{}{
		"Id":          task.Id,
		"GroupId":     task.GroupId,
		"Name":        task.Name,
		"Description": task.Description,
		"Date":        task.Date,
		"StartTime":   task.StartTime,
		"EndTime":     task.EndTime,
		"Location":    task.Utillity,
	}
	database.Update(c, "Tasks", task.PrimaryKey, task.RowKey, taskMap)
	return true
}

// delete a task
func DeleteTask(ctx context.Context, pk string, rk string) bool {
	database.Delete(ctx, "Tasks", pk, rk)
	return true
}

// create a task
func CreateTask(ctx context.Context, task model.TaskDTO) (model.TaskDTO, error) {
	taskMap := map[string]interface{}{
		"Id":          task.Id,
		"GroupId":     task.GroupId,
		"Name":        task.Name,
		"Description": task.Description,
		"Date":        task.Date,
		"StartTime":   task.StartTime,
		"EndTime":     task.EndTime,
		"Location":    task.Utillity,
	}

	err := database.Write(ctx, "Tasks", task.PrimaryKey, task.RowKey, taskMap)
	if err != nil {
		return model.TaskDTO{}, err
	}
	return task, nil
}
