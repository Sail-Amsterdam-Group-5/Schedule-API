package model

import "time"

type TaskDTO struct {
	PrimaryKey  string    `json:"primaryKey"`
	RowKey      string    `json:"rowKey"`
	Id          string    `json:"id"`
	GroupId     string    `json:"groupId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Utillity    string    `json:"utillity"`
}
