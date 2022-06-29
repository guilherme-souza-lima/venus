package repository

import (
	"fmt"
	"gorm.io/gorm"
	"ssVenus/infrastructure/model"
)

type ParticipatingRepository struct {
	db *gorm.DB
}

func NewParticipatingRepository(db *gorm.DB) ParticipatingRepository {
	return ParticipatingRepository{db}
}

func (p ParticipatingRepository) ListParticipating(user, game string) error {
	var entity model.ParticipatingModel
	filter := p.db.Model(&entity).Where("user_id = ? and game_id = ?", user, game)
	if filter.Error != nil {
		return filter.Error
	}
	result, err := filter.Rows()
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}