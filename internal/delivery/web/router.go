package web

import (
	"io/fs"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/seinkytarlicht/remusika/frontend"
)

func RegisteWebRoutes(app *fiber.App) {
	clientFS, err := fs.Sub(frontend.UI, ".output/public")
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/*", static.New("", static.Config{
		FS: clientFS,
		NotFoundHandler: func(c fiber.Ctx) error {
			data, err := fs.ReadFile(clientFS, "index.html")
			if err != nil {
				return c.Status(fiber.StatusNotFound).SendString("index.html not found")
			}
			c.Set("Content-Type", "text/html; charset=utf-8")
			return c.Send(data)
		},
	}))
}
