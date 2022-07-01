package handler

import (
	"github.com/gofiber/fiber/v2"
	"ssVenus/entities"
	"ssVenus/user_case/request"
)

type ParticipatingService interface {
	List(idUser, typeGame, tokenUser string) ([]entities.ParticipatingByUser, error)
	Create(data request.ParticipatingReq) error
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

func (p ParticipatingHandler) CreateParticipating(c *fiber.Ctx) error {
	var data request.ParticipatingReq
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error body parser request. Error: " + err.Error())
	}
	result := p.ParticipatingService.Create(data)
	if result != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error service create participating. Error: " + result.Error())
	}
	return c.Status(fiber.StatusOK).JSON(true)
}
