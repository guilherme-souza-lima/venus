package service

import "ssVenus/entities"

type GamesRepository interface {
	ListGames() ([]entities.ListGames, error)
}

type GamesService struct {
	GamesRepository GamesRepository
}

func NewGamesService(GamesRepository GamesRepository) GamesService {
	return GamesService{GamesRepository}
}

func (g GamesService) ListGames() ([]entities.ListGames, error) {
	return g.GamesRepository.ListGames()
}
