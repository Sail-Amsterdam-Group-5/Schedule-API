package model

type CheckInDTO struct {
	// pk := taskID
	// rk := userID + checkInTime

	PrimaryKey    string `json:"primaryKey"`
	RowKey        string `json:"rowKey"`
	CheckInId     string `json:"checkinId"`
	UserId        string `json:"userId"`
	TaskId        string `json:"taskId"`
	CheckedIn     bool   `json:"checkedIn"`
	CheckInTime   string `json:"checkinTime"`
	CancelledTask bool   `json:"cancelledTask"`
	Reason        string `json:"reason"`
}
