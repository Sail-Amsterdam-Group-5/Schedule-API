package repository

import (
	"context"
	"fmt"
	"schedule-api/database"
	"schedule-api/model"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func SaveCheckIn(ctx context.Context, dto model.CheckInDTO) {
	pk := dto.PrimaryKey
	rk := dto.RowKey

	checkinMap := map[string]interface{}{
		"CheckInId":     dto.CheckInId,
		"UserId":        dto.UserId,
		"TaskId":        dto.TaskId,
		"CheckedIn":     dto.CheckedIn,
		"CheckInTime":   dto.CheckInTime,
		"CancelledTask": dto.CancelledTask,
	}

	database.Write(ctx, "CheckIn", pk, rk, checkinMap)
}

func GetCheckin(ctx context.Context, pk string, rk string) (model.CheckInDTO, error) {
	entity, err := database.ReadSingle(ctx, "CheckIn", pk, rk)
	if err != nil {
		return model.CheckInDTO{}, fmt.Errorf("CheckIn not found")
	}

	dto := model.CheckInDTO{
		PrimaryKey:    entity.Properties["PartitionKey"].(string),
		RowKey:        entity.Properties["RowKey"].(string),
		CheckInId:     entity.Properties["CheckInId"].(string),
		UserId:        entity.Properties["UserId"].(string),
		TaskId:        entity.Properties["TaskId"].(string),
		CheckedIn:     entity.Properties["CheckedIn"].(bool),
		CheckInTime:   entity.Properties["CheckInTime"].(string),
		CancelledTask: entity.Properties["CancelledTask"].(bool),
	}

	return dto, nil
}

func GetAllCheckins(ctx context.Context) ([]aztables.EDMEntity, error) {
	return database.ReadAll(ctx, "CheckIn")
}
func GetAllCheckinDTOs(ctx context.Context) ([]model.CheckInDTO, error) {
	entities, err := database.ReadAll(ctx, "CheckIn")
	if err != nil {
		return nil, err
	}

	var dtos []model.CheckInDTO
	for _, entity := range entities {
		dto := model.CheckInDTO{
			PrimaryKey:    entity.Properties["PartitionKey"].(string),
			RowKey:        entity.Properties["RowKey"].(string),
			CheckInId:     entity.Properties["CheckInId"].(string),
			UserId:        entity.Properties["UserId"].(string),
			TaskId:        entity.Properties["TaskId"].(string),
			CheckedIn:     entity.Properties["CheckedIn"].(bool),
			CheckInTime:   entity.Properties["CheckInTime"].(string),
			CancelledTask: entity.Properties["CancelledTask"].(bool),
		}
		dtos = append(dtos, dto)
	}

	return dtos, nil
}
