package handler

import (
	"github.com/gofiber/fiber/v2"
	"ssVenus/entities"
)

type GamesService interface {
	ListGames() ([]entities.ListGames, error)
}

type GamesHandler struct {
	GamesService GamesService
}

func NewGamesHandler(GamesService GamesService) GamesHandler {
	return GamesHandler{GamesService}
}

func (g GamesHandler) ListGames(c *fiber.Ctx) error {
	list, err := g.GamesService.ListGames()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(list)
}
