package repository

import (
	"schedule-api/database"
	"schedule-api/model"
)

func SaveCheckIn(dto model.CheckInDTO) {
	pk := dto.PrimaryKey
	rk := dto.RowKey

	database.Write("CheckIn", pk, rk, dto)
}

func GetCheckin(id string) {
	return database.ReadSingle("CheckIn", id, id)
}

func GetAllCheckins() {
	return database.ReadAll("CheckIn")
}
