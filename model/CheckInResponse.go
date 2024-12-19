package model

type CheckInResponse struct {
	CheckInId     string `json:"checkinId"`
	UserId        string `json:"userId"`
	TaskId        string `json:"taskId"`
	CheckedIn     bool   `json:"checkedIn"`
	CheckInTime   string `json:"checkinTime"`
	CancelledTask bool   `json:"cancelledTask"`
}
