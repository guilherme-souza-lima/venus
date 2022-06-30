package handler

import (
	"github.com/gofiber/fiber/v2"
	"ssVenus/entities"
)

type ParticipatingService interface {
	List(idUser, typeGame, tokenUser string) ([]entities.ParticipatingByUser, error)
}

type ParticipatingHandler struct {
	ParticipatingService ParticipatingService
}

func NewParticipatingHandler(ParticipatingService ParticipatingService) ParticipatingHandler {
	return ParticipatingHandler{ParticipatingService}
}

func (p ParticipatingHandler) ListParticipating(c *fiber.Ctx) error {
	idUser := c.Params("idUser")
	typeGame := c.Params("typeGame")
	tokenUser := c.Params("tokenUser")
	result, _ := p.ParticipatingService.List(idUser, typeGame, tokenUser)
	return c.Status(fiber.StatusOK).JSON(result)
}
