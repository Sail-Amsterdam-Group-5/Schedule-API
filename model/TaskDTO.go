package model

type TaskDTO struct {
	PrimaryKey  string `json:"primaryKey"` // day + groupID
	RowKey      string `json:"rowKey"`     // startTime + taskID
	Id          string `json:"id"`         // GUID
	GroupId     string `json:"groupId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Location    string `json:"location"` // has to be object
}
