package model

type TaskDTO struct {
	PrimaryKey  string      `json:"primaryKey"` // has to be string
	RowKey      string      `json:"rowKey"`     // has to be string
	Id          int         `json:"id"`
	GroupId     int         `json:"groupId"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Date        string      `json:"date"`
	StartTime   string      `json:"startTime"`
	EndTime     string      `json:"endTime"`
	Location    LocationDTO `json:"location"` // has to be object
}
