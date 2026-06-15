package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/seinkytarlicht/remusika/internal/controller"
)

func MusicRouter(r fiber.Router, c controller.MusicController) {
	r.Post("/reload-folder", c.ReloadFolder)
	r.Get("/get-all", c.GetAll)
	r.Get("/stream/:uuid", c.Stream)
	r.Get("/image/:uuid", c.GetImage)
}
