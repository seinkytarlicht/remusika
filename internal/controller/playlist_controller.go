package controller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/seinkytarlicht/remusika/internal/service"
)

type PlaylistController interface {
	GetAll(c fiber.Ctx) error
}

type PlaylistControllerImpl struct {
	Service service.PlaylistService
}

func (p *PlaylistControllerImpl) GetAll(c fiber.Ctx) error {
	result := p.Service.FindAll(c.Context())

	return c.JSON(result)
}

func NewPlaylistController(service service.PlaylistService) PlaylistController {
	return &PlaylistControllerImpl{
		Service: service,
	}
}
