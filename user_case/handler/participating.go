package handler

import "github.com/gofiber/fiber/v2"

type ParticipatingService interface {
}

type ParticipatingHandler struct {
	ParticipatingService ParticipatingService
}

func NewParticipatingHandler(ParticipatingService ParticipatingService) ParticipatingHandler {
	return ParticipatingHandler{ParticipatingService}
}

func (p ParticipatingHandler) ListParticipating(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("success")
}
