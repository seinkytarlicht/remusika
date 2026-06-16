package controller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/seinkytarlicht/remusika/internal/model"
	"github.com/seinkytarlicht/remusika/internal/service"
)

type PlaylistController interface {
	GetAll(c fiber.Ctx) error
	Create(c fiber.Ctx) error
}

type PlaylistControllerImpl struct {
	Service service.PlaylistService
}

func (p *PlaylistControllerImpl) GetAll(c fiber.Ctx) error {
	result := p.Service.FindAll(c.Context())

	return c.JSON(result)
}

func (p *PlaylistControllerImpl) Create(c fiber.Ctx) error {
	c.Accepts("application/json")

	playlist := new(model.Playlist)

	if err := c.Bind().Body(playlist); err != nil {
		return err
	}

	return c.JSON(playlist)
}

func NewPlaylistController(service service.PlaylistService) PlaylistController {
	return &PlaylistControllerImpl{
		Service: service,
	}
}
