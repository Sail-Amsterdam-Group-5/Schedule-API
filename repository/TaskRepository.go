package repository

import (
	"context"
	"errors"
	"schedule-api/database"
	"schedule-api/model"
)

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
			GroupId:     entity.Properties["GroupId"].(int),
			Name:        entity.Properties["Name"].(string),
			Description: entity.Properties["Description"].(string),
			Date:        entity.Properties["Date"].(string),
			StartTime:   entity.Properties["StartTime"].(string),
			EndTime:     entity.Properties["EndTime"].(string),
			Location:    entity.Properties["Location"].(string),
			// Location: model.LocationDTO{
			// 	Id:          entity.Properties["Location"].(map[string]any)["Id"].(int),
			// 	Name:        entity.Properties["Location"].(map[string]any)["Name"].(string),
			// 	Description: entity.Properties["Location"].(map[string]any)["Description"].(string),
			// 	Address:     entity.Properties["Location"].(map[string]any)["Address"].(string),
			// 	Lat:         entity.Properties["Location"].(map[string]any)["Lat"].(float64),
			// 	Lng:         entity.Properties["Location"].(map[string]any)["Lng"].(float64),
			// },
		}
		tasks = append(tasks, task)
	}

	return tasks
}

// get all for date

func GetAllTaskForDate(ctx context.Context, date string, groupId string) []model.TaskDTO {
	filter := database.BuildDuoFilter("Date", date, "GroupId", groupId)
	entities, err := database.ReadFilter(ctx, "Tasks", filter)
	if err != nil {
		return nil
	}

	var tasks []model.TaskDTO
	for _, entity := range entities {
		task := model.TaskDTO{
			Id:          entity.Properties["Id"].(string),
			GroupId:     entity.Properties["GroupId"].(int),
			Name:        entity.Properties["Name"].(string),
			Description: entity.Properties["Description"].(string),
			Date:        entity.Properties["Date"].(string),
			StartTime:   entity.Properties["StartTime"].(string),
			EndTime:     entity.Properties["EndTime"].(string),
			Location:    entity.Properties["Location"].(string),
			// Location: model.LocationDTO{
			// 	Id:          entity.Properties["Location"].(map[string]any)["Id"].(int),
			// 	Name:        entity.Properties["Location"].(map[string]any)["Name"].(string),
			// 	Description: entity.Properties["Location"].(map[string]any)["Description"].(string),
			// 	Address:     entity.Properties["Location"].(map[string]any)["Address"].(string),
			// 	Lat:         entity.Properties["Location"].(map[string]any)["Lat"].(float64),
			// 	Lng:         entity.Properties["Location"].(map[string]any)["Lng"].(float64),
			// },
		}
		tasks = append(tasks, task)
	}

	return tasks
}

// get by id
func GetTaskById(ctx context.Context, id string) (model.TaskDTO, error) {
	filter := database.BuildFilter("Id", id)
	task, err := database.ReadFilter(ctx, "Tasks", filter)
	if err != nil {
		return model.TaskDTO{}, err
	}

	if len(task) > 0 {
		gid := task[0].Properties["GroupId"].(int32)
		taskDTO := model.TaskDTO{
			PrimaryKey:  task[0].PartitionKey,
			RowKey:      task[0].RowKey,
			Id:          task[0].Properties["Id"].(string),
			GroupId:     int(gid),
			Name:        task[0].Properties["Name"].(string),
			Description: task[0].Properties["Description"].(string),
			Date:        task[0].Properties["Date"].(string),
			StartTime:   task[0].Properties["StartTime"].(string),
			EndTime:     task[0].Properties["EndTime"].(string),
			Location:    task[0].Properties["Location"].(string),
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
		"Location":    task.Location,
		// "Location": map[string]interface{}{
		// 	"Id":          task.Location.Id,
		// 	"Name":        task.Location.Name,
		// 	"Description": task.Location.Description,
		// 	"Address":     task.Location.Address,
		// 	"Lat":         task.Location.Lat,
		// 	"Lng":         task.Location.Lng,
		// },
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
		"Location":    task.Location,
		// "Location": map[string]interface{}{
		// 	"Id": task.Location,
		// "Name":        task.Location.Name,
		// "Description": task.Location.Description,
		// "Address":     task.Location.Address,
		// "Lat":         task.Location.Lat,
		// "Lng":         task.Location.Lng,
	}

	err := database.Write(ctx, "Tasks", task.PrimaryKey, task.RowKey, taskMap)
	if err != nil {
		return model.TaskDTO{}, err
	}
	return task, nil
}
