package model

import (
	"time"
)

type LocationDTO struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	Location  Location  `json:"location"`
	Ocean     string    `json:"ocean"`
}
