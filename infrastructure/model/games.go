package model

import (
	"ssVenus/entities"
	"time"
)

type GamesModel struct {
	ID            string    `gorm:"PrimaryKey"`
	Name          string    `gorm:"column:name"`
	FormattedDate string    `gorm:"column:formatted_date"`
	StartData     time.Time `gorm:"column:start_date"`
	EndDate       time.Time `gorm:"column:end_date"`
	EndGame       time.Time `gorm:"column:end_game"`
	State         int       `gorm:"column:state"`
	Photo         string    `gorm:"column:photo"`
	TypeGame      int       `gorm:"column:type_game"`
	PointsGG      int       `gorm:"column:points_gg"`
}

func (ref *GamesModel) TableName() string {
	return "games"
}

func (ref GamesModel) ToDomain() entities.ListGames {
	return entities.ListGames{
		ID:            ref.ID,
		Name:          ref.Name,
		FormattedDate: ref.FormattedDate,
		StartData:     ref.StartData,
		EndDate:       ref.EndDate,
		EndGame:       ref.EndGame,
		State:         ref.State,
		Photo:         ref.Photo,
		TypeGame:      ref.TypeGame,
		PointsGG:      ref.PointsGG,
	}
}
