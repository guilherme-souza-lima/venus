package repository

import (
	"gorm.io/gorm"
	"ssVenus/entities"
	"ssVenus/infrastructure/model"
)

type GamesRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) GamesRepository {
	return GamesRepository{db}
}

func (g GamesRepository) ListGames() ([]entities.ListGames, error) {
	var entity model.GamesModel
	filter := g.db.Model(&entity)
	if filter.Error != nil {
		return nil, filter.Error
	}
	result, err := filter.Rows()
	if err != nil {
		return nil, err
	}

	var list []entities.ListGames
	for result.Next() {
		var row model.GamesModel
		result.Scan(
			&row.ID,
			&row.UUID,
			&row.Name,
			&row.FormattedDate,
			&row.StartData,
			&row.EndDate,
			&row.EndGame,
			&row.State,
			&row.Photo,
		)
		list = append(list, row.ToDomain())
	}

	return list, nil
}
