package model

import (
	"github.com/gofrs/uuid"
	"ssVenus/entities"
)

type ParticipatingModel struct {
	ID   int    `gorm:"PrimaryKey"`
	UUID string `gorm:"column:uuid"`
	User string `gorm:"column:user_id"`
	Game string `gorm:"column:game_id"`
	Nick string `gorm:"column:nick"`
}

func (ref *ParticipatingModel) FromDomain(data entities.Participating) {
	uuidGenerator, _ := uuid.NewV4()
	ref.UUID = uuidGenerator.String()
	ref.User = data.UserID
	ref.Game = data.GameID
	ref.Nick = data.Nick
}

func (ref *ParticipatingModel) TableName() string {
	return "participating"
}
