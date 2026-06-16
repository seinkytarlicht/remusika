package http

import (
	"github.com/gofiber/fiber/v3"
)

func RegisterApiRoutes(app *fiber.App, musicC MusicController, playlistC PlaylistController, shutdownC ShutdownController) {
	api := app.Group("/api")

	musicR := api.Group("/music")
	musicR.Post("/reload-folder", musicC.ReloadFolder)
	musicR.Get("/get-all", musicC.GetAll)
	musicR.Get("/stream/:uuid", musicC.Stream)
	musicR.Get("/image/:uuid", musicC.GetImage)

	playlistR := api.Group("/playlist")
	playlistR.Get("/get-all", playlistC.GetAll)
	playlistR.Post("/create", playlistC.Create)

	shutdownR := api.Group("/shutdown")
	shutdownR.Post("/off", shutdownC.Off)
}
