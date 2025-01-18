package repository

import (
	"context"
	"fmt"
	"schedule-api/database"
	"schedule-api/model"
)

func SaveCheckIn(ctx context.Context, dto model.CheckInDTO) error {
	pk := dto.PrimaryKey
	rk := dto.RowKey

	checkinMap := map[string]interface{}{
		"CheckInId":     dto.CheckInId,
		"UserId":        dto.UserId,
		"TaskId":        dto.TaskId,
		"CheckedIn":     dto.CheckedIn,
		"CheckInTime":   dto.CheckInTime,
		"CancelledTask": dto.CancelledTask,
		"Reason":        dto.Reason,
	}

	return database.Write(ctx, "CheckIn", pk, rk, checkinMap)
}

func GetCheckin(ctx context.Context, userId string, taskId string) (model.CheckInDTO, error) {
	filter := database.BuildDuoFilter("UserId", userId, "TaskId", taskId)
	entities, err := database.ReadFilter(ctx, "CheckIn", filter)
	if err != nil || len(entities) == 0 {
		return model.CheckInDTO{}, fmt.Errorf("CheckIn not found")
	}

	entity := entities[0]
	dto := model.CheckInDTO{
		PrimaryKey:    entity.PartitionKey,
		RowKey:        entity.RowKey,
		CheckInId:     entity.Properties["CheckInId"].(string),
		UserId:        entity.Properties["UserId"].(string),
		TaskId:        entity.Properties["TaskId"].(string),
		CheckedIn:     entity.Properties["CheckedIn"].(bool),
		CheckInTime:   parseTime(entity.Properties["CheckInTime"].(string)),
		CancelledTask: entity.Properties["CancelledTask"].(bool),
		Reason:        entity.Properties["Reason"].(string),
	}

	return dto, nil
}

func GetAllCheckins(ctx context.Context) ([]model.CheckInDTO, error) {
	entities, err := database.ReadAll(ctx, "CheckIn")
	if err != nil {
		return nil, err
	}

	var dtos []model.CheckInDTO
	for _, entity := range entities {
		dto := model.CheckInDTO{
			PrimaryKey:    entity.PartitionKey,
			RowKey:        entity.RowKey,
			CheckInId:     entity.Properties["CheckInId"].(string),
			UserId:        entity.Properties["UserId"].(string),
			TaskId:        entity.Properties["TaskId"].(string),
			CheckedIn:     entity.Properties["CheckedIn"].(bool),
			CheckInTime:   parseTime(entity.Properties["CheckInTime"].(string)),
			CancelledTask: entity.Properties["CancelledTask"].(bool),
			Reason:        entity.Properties["Reason"].(string),
		}
		dtos = append(dtos, dto)
	}
	return dtos, nil
}

func CheckCheckin(ctx context.Context, userId string, taskId string) (bool, error) {
	filter := database.BuildDuoFilter("UserId", userId, "TaskId", taskId)
	entities, err := database.ReadFilter(ctx, "CheckIn", filter)
	if err != nil {
		return false, fmt.Errorf("CheckIn not found")
	}

	for _, entity := range entities {
		dto := model.CheckInDTO{
			PrimaryKey:    entity.PartitionKey,
			RowKey:        entity.RowKey,
			CheckInId:     entity.Properties["CheckInId"].(string),
			UserId:        entity.Properties["UserId"].(string),
			TaskId:        entity.Properties["TaskId"].(string),
			CheckedIn:     entity.Properties["CheckedIn"].(bool),
			CheckInTime:   parseTime(entity.Properties["CheckInTime"].(string)),
			CancelledTask: entity.Properties["CancelledTask"].(bool),
			Reason:        entity.Properties["Reason"].(string),
		}
		if dto.CheckedIn && !dto.CancelledTask {
			return true, nil
		}
	}
	return false, nil
}

func UpdateCheckin(ctx context.Context, dto model.CheckInDTO) error {
	pk := dto.PrimaryKey
	rk := dto.RowKey

	checkinMap := map[string]interface{}{
		"CheckInId":     dto.CheckInId,
		"UserId":        dto.UserId,
		"TaskId":        dto.TaskId,
		"CheckedIn":     dto.CheckedIn,
		"CheckInTime":   dto.CheckInTime,
		"CancelledTask": dto.CancelledTask,
		"Reason":        dto.Reason,
	}

	return database.Update(ctx, "CheckIn", pk, rk, checkinMap)
}
