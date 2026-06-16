package web

import (
	"io/fs"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/seinkytarlicht/remusika/frontend"
	"github.com/seinkytarlicht/remusika/helper"
)

func RegisteWebRoutes(app *fiber.App) {
	go helper.StartTray(app)
	go helper.OpenBrowser()

	clientFS, err := fs.Sub(frontend.UI, ".output/public")
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/*", static.New("", static.Config{FS: clientFS}))
	app.Get("*", static.New("index.html", static.Config{FS: clientFS}))
}
