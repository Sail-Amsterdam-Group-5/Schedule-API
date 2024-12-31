package service

import (
	"context"
	"schedule-api/model"
	"schedule-api/repository"
	"time"

	"github.com/google/uuid"
)

func Checkin(ctx context.Context, userId string, taskId string) model.CheckInResponse {

	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	dto := model.CheckInDTO{
		PrimaryKey:    taskId,
		RowKey:        userId + time.Now().Format("HH:MM"),
		CheckInId:     id.String(),
		UserId:        userId,
		TaskId:        taskId,
		CheckedIn:     true,
		CheckInTime:   time.Now().Format("HH:MM"),
		CancelledTask: false,
	}

	repository.SaveCheckIn(ctx, dto)
	return model.CheckInResponse{CheckInId: dto.CheckInId, UserId: dto.UserId, TaskId: dto.TaskId, CheckedIn: dto.CheckedIn, CheckInTime: dto.CheckInTime, CancelledTask: dto.CancelledTask}
}

func CancelTask(ctx context.Context, userId string, taskId string) model.CheckInResponse {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	dto := model.CheckInDTO{
		PrimaryKey:    taskId,
		RowKey:        userId + time.Now().Format("HH:MM"),
		CheckInId:     id.String(),
		UserId:        userId,
		TaskId:        taskId,
		CheckedIn:     false,
		CheckInTime:   time.Now().Format("HH:MM"),
		CancelledTask: true,
	}
	repository.SaveCheckIn(ctx, dto)
	return model.CheckInResponse{CheckInId: dto.CheckInId, UserId: dto.UserId, TaskId: dto.TaskId, CheckedIn: dto.CheckedIn, CheckInTime: dto.CheckInTime, CancelledTask: dto.CancelledTask}
}
