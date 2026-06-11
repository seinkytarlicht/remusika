package helper

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v3"
)

var shutdownOnce sync.Once

func Shutdown(app *fiber.App) {
	shutdownOnce.Do(func() {
		CleanUpTemp()

		if err := app.Shutdown(); err != nil {
			log.Println(err)
		}
	})
}
