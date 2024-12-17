package service

import (
	"schedule-api/model"
	"schedule-api/repository"
	"time"

	"github.com/google/uuid"
)

func Checkin(userId string, taskId string) model.CheckInResponce {

	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	dto := model.CheckInDTO{
		CheckInId:     id.String(),
		UserId:        userId,
		TaskId:        taskId,
		CheckedIn:     true,
		CheckInTime:   time.Now().Format("HH:MM"),
		CancelledTask: false,
	}

	repository.SaveCheckIn(dto)
	return dto
}
