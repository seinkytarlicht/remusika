package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/seinkytarlicht/remusika/internal/controller"
)

func ShutdownRouter(r fiber.Router) {
	shutdownController := controller.NewShutdownController()

	r.Post("/off", shutdownController.Off)

}
