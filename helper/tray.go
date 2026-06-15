package helper

import (
	_ "embed"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gogpu/systray"
	"github.com/seinkytarlicht/remusika/assets"
)

func StartTray(app *fiber.App) {

	tray := systray.New()

	menu := systray.NewMenu()
	menu.Add("Show", func() {
		OpenBrowser()
	})
	menu.Add("Exit", func() {
		Shutdown(app)
	})

	tray.SetIcon(assets.IconTray).
		SetTooltip("ReMusika").
		SetMenu(menu)

	tray.OnClick(func() {
		OpenBrowser()
	})

	tray.Show()
	if err := tray.Run(); err != nil {
		fmt.Println("error:", err)
	}
}
