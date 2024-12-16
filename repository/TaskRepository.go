package repository

import "schedule-api/model"

// get all for user
func GetAllTaskForUser(userId string) []model.TaskDTO {

	// the connection to the database is established here

	// the query is executed here

	//return db.Query("SELECT * FROM tasks WHERE user_id = ?", userId)

	return []model.TaskDTO{
		{
			Id:          1,
			GroupId:     1,
			Name:        "Task 1",
			Description: "Description 1",
			Date:        "2019-01-01",
			StartTime:   "09:00",
			EndTime:     "10:00",
			Location: model.LocationDTO{
				Id:          1,
				Name:        "Location 1",
				Description: "Description 1",
				Address:     "Address 1",
				Lat:         1.0,
				Lng:         1.0,
			},
		},
	}
	// 	var tasks []model.TaskDTO

	// 	// Simulate database query and error handling
	// 	err := db.Query("SELECT * FROM tasks WHERE user_id = ?", userId).Scan(&tasks)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return tasks, nil
	// }
}

// get all for date

// get by id

// update a task

// delete a task

// create a task
