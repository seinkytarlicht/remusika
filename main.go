package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/seinkytarlicht/remusika/config"
	"github.com/seinkytarlicht/remusika/helper"
	"github.com/seinkytarlicht/remusika/internal/delivery/http"
)

//go:embed frontend/.output/public/*
var clientUI embed.FS

//go:embed assets/logo.png
var iconTray []byte

func main() {
	// the holy spaghetti, i love spaghetti
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	var serverOnly bool
	flag.BoolVar(&serverOnly, "server", false, "No client, server only")
	flag.BoolVar(&serverOnly, "s", false, "No client, server only (Short)")
	flag.Parse()

	app := fiber.New()
	app.Use(logger.New())
	app.Route("/api", http.SetupApiRoute)
	if !serverOnly {
		_ClientSetup(app)
	}
	go func() {
		<-sig

		helper.Shutdown(app)
	}()
	app.Listen(config.ServerAddr)

}

func _ClientSetup(app *fiber.App) {
	go helper.StartTray(app, iconTray)
	go helper.OpenBrowser()

	clientFS, err := fs.Sub(clientUI, "frontend/.output/public")
	if err != nil {
		log.Fatal(err)
	}
	app.Get("/*", static.New("", static.Config{
		FS: clientFS,
	}))

	// Fallback not found
	app.Use(func(c fiber.Ctx) error {
		c.Path("/")

		return static.New("", static.Config{
			FS: clientFS,
		})(c)
	})
}
