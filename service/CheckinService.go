package service

import (
	"context"
	"schedule-api/model"
	"schedule-api/repository"
	"time"

	"github.com/google/uuid"
)

func Checkin(ctx context.Context, userId string, taskId string) (model.CheckInResponse, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	formatedTime := time.Now().Format("15:04")
	dto := model.CheckInDTO{
		PrimaryKey:    taskId,
		RowKey:        userId + time.Now().Format(time.Stamp),
		CheckInId:     id.String(),
		UserId:        userId,
		TaskId:        taskId,
		CheckedIn:     true,
		CheckInTime:   formatedTime,
		CancelledTask: false,
	}

	err = repository.SaveCheckIn(ctx, dto)
	if err != nil {
		return model.CheckInResponse{}, err
	}

	response := model.CheckInResponse{
		CheckInId:     dto.CheckInId,
		UserId:        dto.UserId,
		TaskId:        dto.TaskId,
		CheckedIn:     dto.CheckedIn,
		CheckInTime:   dto.CheckInTime,
		CancelledTask: dto.CancelledTask}

	return response, nil
}

func CancelTask(ctx context.Context, userId string, taskId string) (model.CheckInResponse, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	formatedTime := time.Now().Format("15:04")
	dto := model.CheckInDTO{
		PrimaryKey:    taskId,
		RowKey:        userId + time.Now().Format(time.Stamp),
		CheckInId:     id.String(),
		UserId:        userId,
		TaskId:        taskId,
		CheckedIn:     false,
		CheckInTime:   formatedTime,
		CancelledTask: true,
	}
	err = repository.SaveCheckIn(ctx, dto)
	if err != nil {
		return model.CheckInResponse{}, err
	}

	response := model.CheckInResponse{
		CheckInId:     dto.CheckInId,
		UserId:        dto.UserId,
		TaskId:        dto.TaskId,
		CheckedIn:     dto.CheckedIn,
		CheckInTime:   dto.CheckInTime,
		CancelledTask: dto.CancelledTask,
	}
	return response, nil
}

func GetAllCheckIns(ctx context.Context) ([]model.CheckInResponse, error) {
	checkIns, err := repository.GetAllCheckins(ctx)
	if err != nil {
		return nil, err
	}

	var response []model.CheckInResponse
	for _, checkIn := range checkIns {
		checkInResponse := model.CheckInResponse{
			CheckInId:     checkIn.CheckInId,
			UserId:        checkIn.UserId,
			TaskId:        checkIn.TaskId,
			CheckedIn:     checkIn.CheckedIn,
			CheckInTime:   checkIn.CheckInTime,
			CancelledTask: checkIn.CancelledTask,
		}
		response = append(response, checkInResponse)
	}
	return response, nil
}
