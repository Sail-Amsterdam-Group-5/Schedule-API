package model

import "time"

type CheckInDTO struct {
	PrimaryKey    string    `json:"primaryKey"`
	RowKey        string    `json:"rowKey"`
	CheckInId     string    `json:"checkinId"`
	UserId        string    `json:"userId"`
	TaskId        string    `json:"taskId"`
	CheckedIn     bool      `json:"checkedIn"`
	CheckInTime   time.Time `json:"checkinTime"`
	CancelledTask bool      `json:"cancelledTask"`
	Reason        string    `json:"reason"`
}
