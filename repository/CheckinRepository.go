package repository

import (
	"schedule-api/model"
	"schedule-api/repository"
)

func SaveCheckIn(dto model.CheckInDTO) {
	pk := dto.TaskId
	rk := dto.UserId + dto.CheckInTime

	repository.Write("CheckIn", pk, rk, dto)
}
