package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/seinkytarlicht/remusika/internal/model"
	"github.com/seinkytarlicht/remusika/internal/service"
)

type PlaylistController interface {
	GetAll(c fiber.Ctx) error
	Create(c fiber.Ctx) error
	Update(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
	AddMusic(c fiber.Ctx) error
	RemoveMusic(c fiber.Ctx) error
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

	res := p.Service.Create(c.Context(), *playlist)

	return c.Status(201).JSON(res)
}

func (p *PlaylistControllerImpl) Update(c fiber.Ctx) error {
	c.Accepts("application/json")

	playlist := new(model.Playlist)

	if err := c.Bind().Body(playlist); err != nil {
		return err
	}

	res := p.Service.Update(c.Context(), *playlist)

	return c.JSON(res)
}

func (p *PlaylistControllerImpl) Delete(c fiber.Ctx) error {
	id := fiber.Params[int64](c, "id")

	res := p.Service.Delete(c.Context(), id)

	if !res {
		return c.Status(400).JSON(fiber.Map{
			"messages": "Delete Failed",
		})
	}

	return c.JSON(fiber.Map{
		"messages": "Delete Success",
	})
}

func (p *PlaylistControllerImpl) AddMusic(c fiber.Ctx) error {
	panic("unimplemented")
}

func (p *PlaylistControllerImpl) RemoveMusic(c fiber.Ctx) error {
	panic("unimplemented")
}

func NewPlaylistController(service service.PlaylistService) PlaylistController {
	return &PlaylistControllerImpl{
		Service: service,
	}
}
