package model

type Task struct {
	Id          string      `json:"id"`
	GroupId     string      `json:"groupId"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Date        string      `json:"date"`
	StartTime   string      `json:"startTime"`
	EndTime     string      `json:"endTime"`
	Location    LocationDTO `json:"location"` //has to be object
}
