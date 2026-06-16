package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/seinkytarlicht/remusika/config"
	"github.com/seinkytarlicht/remusika/helper"
	"github.com/seinkytarlicht/remusika/internal/database"
	apiDelivery "github.com/seinkytarlicht/remusika/internal/delivery/http"
	webDelivery "github.com/seinkytarlicht/remusika/internal/delivery/web"
	"github.com/seinkytarlicht/remusika/internal/repository"
	"github.com/seinkytarlicht/remusika/internal/service"
)

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
	runApiServer(app)

	if !serverOnly {
		runWebsiteServer(app)
	}

	// Get the list of routes
	routes := app.GetRoutes()

	// Pretty-print the routes as JSON
	routesJSON, err := json.MarshalIndent(routes, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(routesJSON))

	go func() {
		<-sig

		helper.Shutdown(app)
	}()
	app.Listen(config.ServerAddr)
}

func runApiServer(app *fiber.App) {
	db, err := database.New()

	if err != nil {
		panic(err)
	}

	musicMemo := repository.NewMusicMemo()
	musicService := service.NewMusicService(musicMemo)
	musicController := apiDelivery.NewMusicController(musicService)

	playlistRepo := repository.NewPlaylistRepository(db)
	playlistService := service.NewPlaylistService(playlistRepo)
	playlistController := apiDelivery.NewPlaylistController(playlistService)

	shutdownController := apiDelivery.NewShutdownController()

	apiDelivery.RegisterApiRoutes(app, musicController, playlistController, shutdownController)
}

func runWebsiteServer(app *fiber.App) {
	webDelivery.RegisteWebRoutes(app)
}
