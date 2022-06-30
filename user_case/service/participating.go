package service

import "ssVenus/entities"

type ParticipatingRepository interface {
	ListParticipating(UserID, TypeGame string) (list []entities.ParticipatingByUser, err error)
}

type ParticipatingService struct {
	ParticipatingRepository ParticipatingRepository
}

func NewParticipatingService(ParticipatingRepository ParticipatingRepository) ParticipatingService {
	return ParticipatingService{ParticipatingRepository}
}

func (p ParticipatingService) List(idUser, typeGame, tokenUser string) ([]entities.ParticipatingByUser, error) {
	return p.ParticipatingRepository.ListParticipating(idUser, typeGame)
}
