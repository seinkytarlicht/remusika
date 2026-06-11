package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/seinkytarlicht/remusika/internal/controller"
	"github.com/seinkytarlicht/remusika/internal/repository"
	"github.com/seinkytarlicht/remusika/internal/service"
)

func MusicRouter(r fiber.Router) {
	musicController := controller.NewMusicController(service.NewMusicService(repository.NewMusicMemo()))

	r.Post("/reload-folder", musicController.ReloadFolder)
	r.Get("/get-all", musicController.GetAll)
	r.Get("/stream/:uuid", musicController.Stream)
	r.Get("/metadata/:uuid", musicController.GetMetadata)
	r.Get("/image/:uuid", musicController.GetImage)
}
