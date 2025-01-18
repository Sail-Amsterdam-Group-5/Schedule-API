package model

import "time"

type CheckInResponse struct {
	CheckInId     string    `json:"checkinId"`
	UserId        string    `json:"userId"`
	TaskId        string    `json:"taskId"`
	CheckedIn     bool      `json:"checkedIn"`
	CheckInTime   time.Time `json:"checkinTime"`
	CancelledTask bool      `json:"cancelledTask"`
	Reason        string    `json:"reason"`
}
