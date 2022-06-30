package repository

import (
	"gorm.io/gorm"
	"ssVenus/entities"
	"ssVenus/infrastructure/model"
)

type ParticipatingRepository struct {
	db *gorm.DB
}

func NewParticipatingRepository(db *gorm.DB) ParticipatingRepository {
	return ParticipatingRepository{db}
}

func (p ParticipatingRepository) ListParticipating(UserID, TypeGame string) (list []entities.ParticipatingByUser, err error) {
	var entity model.GamesModel
	filter := p.db.Model(&entity).
		Select("games.id,games.name,games.formatted_date,games.start_date,games.end_date,"+
			"games.end_game,games.state,games.photo,games.type_game,games.points_gg,"+
			"participating.id as participating_id,participating.nick").
		Joins("left join participating on participating.game_id = games.id and participating.user_id = ?", UserID).
		Where("games.type_game = ?", TypeGame)

	result, _ := filter.Rows()

	for result.Next() {
		var row entities.ParticipatingByUser
		result.Scan(
			&row.ID,
		)
		list = append(list, row)
	}
	return list, nil
}
