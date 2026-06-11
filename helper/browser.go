package helper

import (
	"os/exec"
	"runtime"

	"github.com/seinkytarlicht/remusika/config"
)

func OpenBrowser() {
	var cmd string
	var args []string
	url := config.Url

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // Linux
		cmd = "xdg-open"
		args = []string{url}
	}

	_ = exec.Command(cmd, args...).Start()
}
