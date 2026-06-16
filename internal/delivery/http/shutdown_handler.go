package http

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/seinkytarlicht/remusika/helper"
)

type ShutdownController interface {
	Off(c fiber.Ctx) error
}

type ShutdownControllerImpl struct {
}

func (s *ShutdownControllerImpl) Off(c fiber.Ctx) error {
	go func() {
		time.Sleep(time.Duration(100))
		go helper.Shutdown(c.App())
	}()
	return c.SendString("Shutdown")
}

func NewShutdownController() ShutdownController {
	return &ShutdownControllerImpl{}
}
