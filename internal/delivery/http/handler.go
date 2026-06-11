package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/seinkytarlicht/remusika/internal/delivery/route"
)

func SetupApiRoute(r fiber.Router) {
	r.Route("/music", route.MusicRouter)
	r.Route("/shutdown", route.ShutdownRouter)
}
