package config

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	ServerPort = ":41991"
	ServerAddr = "127.0.0.1" + ServerPort
	AppUrl     = "http://" + ServerAddr
	AppName    = "ReMusika"
	AppVersion = "v0.1.1-beta.1"
)

func GetAppFolder() string {
	configDir, _ := os.UserConfigDir()

	appPath := filepath.Join(configDir, strings.ToLower(AppName))

	return appPath
}
