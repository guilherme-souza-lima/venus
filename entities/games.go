package entities

import (
	"time"
)

type ListGames struct {
	ID            string    `json:"id"`
	UUID          string    `json:"uuid"`
	Name          string    `json:"name"`
	FormattedDate string    `json:"formatted_date"`
	StartData     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	EndGame       time.Time `json:"end_game"`
	State         int       `json:"state"`
	Photo         string    `json:"photo"`
}
