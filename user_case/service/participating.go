package service

import (
	"ssVenus/entities"
	"ssVenus/user_case/request"
)

type ParticipatingRepository interface {
	ListParticipating(UserID, TypeGame string) (list []entities.ParticipatingByUser, err error)
	CreateParticipating(data entities.Participating) error
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

func (p ParticipatingService) Create(data request.ParticipatingReq) error {
	participating := entities.Participating{
		UserID: data.UserID,
		GameID: data.GameID,
		Nick:   data.Nick,
	}
	return p.ParticipatingRepository.CreateParticipating(participating)
}
