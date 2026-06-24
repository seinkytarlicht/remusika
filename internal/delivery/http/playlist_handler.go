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
		return c.Status(400).JSON(fiber.Map{
			"messages": err.Error(),
		})
	}

	if playlist.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"messages": "No Body",
		})
	}

	res, err := p.Service.Create(c.Context(), *playlist)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"messages": err.Error(),
		})
	}

	return c.Status(201).JSON(res)
}

func (p *PlaylistControllerImpl) Update(c fiber.Ctx) error {
	c.Accepts("application/json")

	playlist := new(model.Playlist)

	if err := c.Bind().Body(playlist); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"messages": err.Error(),
		})
	}

	if playlist.Id == 0 || playlist.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"messages": "No Body",
		})
	}

	res, err := p.Service.Update(c.Context(), *playlist)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"messages": err.Error(),
		})
	}

	return c.JSON(res)
}

func (p *PlaylistControllerImpl) Delete(c fiber.Ctx) error {
	id := fiber.Params[uint64](c, "id")

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
	c.Accepts("application/json")

	playlistItem := new(model.PlaylistItem)

	if err := c.Bind().Body(playlistItem); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"messages": err.Error(),
		})
	}

	if playlistItem.PlaylistId == 0 || playlistItem.MusicId == "" {
		return c.Status(400).JSON(fiber.Map{
			"messages": "No Body",
		})
	}

	err := p.Service.AddMusic(c.Context(), playlistItem.PlaylistId, playlistItem.MusicId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"messages": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"messages": "Added music to playlist",
	})
}

func (p *PlaylistControllerImpl) RemoveMusic(c fiber.Ctx) error {
	c.Accepts("application/json")

	playlistItem := new(model.PlaylistItem)

	if err := c.Bind().Body(playlistItem); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"messages": err.Error(),
		})
	}

	if playlistItem.Id == 0 {
		return c.Status(400).JSON(fiber.Map{
			"messages": "No Body",
		})
	}

	err := p.Service.RemoveMusic(c.Context(), playlistItem.Id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"messages": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"messages": "Music removed from playlist",
	})
}

func NewPlaylistController(service service.PlaylistService) PlaylistController {
	return &PlaylistControllerImpl{
		Service: service,
	}
}
