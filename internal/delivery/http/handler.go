package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/seinkytarlicht/remusika/internal/controller"
	"github.com/seinkytarlicht/remusika/internal/database"
	"github.com/seinkytarlicht/remusika/internal/delivery/route"
	"github.com/seinkytarlicht/remusika/internal/repository"
	"github.com/seinkytarlicht/remusika/internal/service"
)

func SetupApiRoute(r fiber.Router) {
	db, err := database.New()

	if err != nil {
		panic(err)
	}

	musicMemo := repository.NewMusicMemo()
	musicService := service.NewMusicService(musicMemo)
	musicController := controller.NewMusicController(musicService)

	playlistRepo := repository.NewPlaylistRepository(db)
	playlistService := service.NewPlaylistService(playlistRepo)
	playlistController := controller.NewPlaylistController(playlistService)

	r.Route("/music", (func(r fiber.Router) { route.MusicRouter(r, musicController) }))
	r.Route("/playlist", (func(r fiber.Router) { route.PlaylistRouter(r, playlistController) }))
	r.Route("/shutdown", route.ShutdownRouter)
}
