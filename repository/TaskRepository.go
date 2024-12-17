package repository

import (
	"schedule-api/model"
	"schedule-api/repository"
)

// get all for user
func GetAllTaskForUser(groupId string) []model.TaskDTO {

	// return []model.TaskDTO{
	// 	{
	// 		Id:          1,
	// 		GroupId:     1,
	// 		Name:        "Task 1",
	// 		Description: "Description 1",
	// 		Date:        "2019-01-01",
	// 		StartTime:   "09:00",
	// 		EndTime:     "10:00",
	// 		Location: model.LocationDTO{
	// 			Id:          1,
	// 			Name:        "Location 1",
	// 			Description: "Description 1",
	// 			Address:     "Address 1",
	// 			Lat:         1.0,
	// 			Lng:         1.0,
	// 		},
	// 	},
	// }

	return repository.ReadFilter("Tasks", "GroupId", groupId)
}

// get all for date

func GetAllTaskForDate(date string, groupId string) []model.TaskDTO {

	tasks := repository.ReadDuoFilter("Tasks", "Date", date, "GroupId", groupId)

	return tasks

}

// get by id

func GetTaskById(id string) model.TaskDTO {

	task := repository.ReadSingle("Tasks", "Id", id)

	return task

}

// update a task

func UpdateTask(task model.TaskDTO) bool {
	repository.Update("Tasks", task.PrimaryKey, task.RowKey, task)
	return true
}

// delete a task

func DeleteTask(pk string, rk string) bool {
	repository.Delete("Tasks", pk, rk)
	return true
}

// create a task

func CreateTask(task model.TaskDTO) {
	repository.Create("Tasks", task)
}
