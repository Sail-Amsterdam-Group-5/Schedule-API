package service

import (
	"context"
	"schedule-api/model"
	"schedule-api/repository"
	"time"

	"github.com/go-errors/errors"
	"github.com/google/uuid"
)

func Checkin(ctx context.Context, userId string, taskId string) (model.CheckInResponse, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	now := time.Now()
	dto := model.CheckInDTO{
		PrimaryKey:    taskId,
		RowKey:        userId + now.GoString(),
		CheckInId:     id.String(),
		UserId:        userId,
		TaskId:        taskId,
		CheckedIn:     true,
		CheckInTime:   now,
		CancelledTask: false,
		Reason:        "",
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

func CancelTask(ctx context.Context, userId string, taskId string, reason string) (model.CheckInResponse, error) {
	existingCheckIn, err := repository.GetCheckin(ctx, userId, taskId) // doet het niet
	if existingCheckIn.Reason != "" {
		return model.CheckInResponse{}, errors.New("Task already cancelled")
	}
	if err == nil && existingCheckIn.Reason != "" {
		existingCheckIn.CancelledTask = true
		existingCheckIn.Reason = reason
		err = repository.UpdateCheckin(ctx, existingCheckIn)
		if err != nil {
			return model.CheckInResponse{}, err
		}

		response := model.CheckInResponse{
			CheckInId:     existingCheckIn.CheckInId,
			UserId:        existingCheckIn.UserId,
			TaskId:        existingCheckIn.TaskId,
			CheckedIn:     existingCheckIn.CheckedIn,
			CheckInTime:   existingCheckIn.CheckInTime,
			CancelledTask: existingCheckIn.CancelledTask,
			Reason:        existingCheckIn.Reason,
		}
		return response, nil
	}

	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	now := time.Now()
	dto := model.CheckInDTO{
		PrimaryKey:    taskId,
		RowKey:        userId + id.String(),
		CheckInId:     id.String(),
		UserId:        userId,
		TaskId:        taskId,
		CheckedIn:     false,
		CheckInTime:   now,
		CancelledTask: true,
		Reason:        reason,
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
		Reason:        dto.Reason,
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

func GetCheckInForTask(ctx context.Context, userId string, taskId string) (bool, error) {
	checkIn, err := repository.CheckCheckin(ctx, userId, taskId)
	if err != nil {
		return false, err
	}
	return checkIn, nil
}
