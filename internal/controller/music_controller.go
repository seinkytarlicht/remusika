package controller

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/seinkytarlicht/remusika/internal/service"
)

type MusicController interface {
	GetAll(c fiber.Ctx) error
	Stream(c fiber.Ctx) error
	GetImage(c fiber.Ctx) error
	ReloadFolder(c fiber.Ctx) error
}

type MusicControllerImpl struct {
	Service service.MusicService
}

func (m *MusicControllerImpl) GetAll(c fiber.Ctx) error {
	music := m.Service.GetAll()

	return c.JSON(music)
}

func (m *MusicControllerImpl) Stream(c fiber.Ctx) error {
	uuid := c.Params("uuid", "")

	music, err := m.Service.Get(uuid)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
			"uuid":    uuid,
		})
	}

	c.Set("Accept-Ranges", "bytes")
	return c.SendFile(music.Path)
}

func (m *MusicControllerImpl) GetImage(c fiber.Ctx) error {
	uuid := c.Params("uuid", "")

	musicPath, err := m.Service.GetImage(uuid)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": err.Error(),
			"uuid":    uuid,
		})
	}

	imageBytes, err := os.ReadFile(musicPath)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": err.Error(),
			"uuid":    uuid,
		})
	}

	contentType := http.DetectContentType(imageBytes)

	c.Set("Content-Type", contentType)
	return c.Send(imageBytes)
}

func (m *MusicControllerImpl) ReloadFolder(c fiber.Ctx) error {
	m.Service.ReloadFolder()

	return c.JSON(fiber.Map{
		"message": "Folder Reloaded",
	})
}

func NewMusicController(s service.MusicService) MusicController {
	return &MusicControllerImpl{
		Service: s,
	}
}
