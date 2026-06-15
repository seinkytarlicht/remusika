package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/seinkytarlicht/remusika/internal/controller"
)

func PlaylistRouter(r fiber.Router, c controller.PlaylistController) {
	r.Get("/get-all", c.GetAll)
}
