package model

import (
	"time"
)

type Task struct {
	Id          string      `json:"id"`
	GroupId     string      `json:"groupId"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Date        time.Time   `json:"date"`
	StartTime   time.Time   `json:"startTime"`
	EndTime     time.Time   `json:"endTime"`
	Location    LocationDTO `json:"location"`
}
