package config

import (
	"errors"
	"io/fs"
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

	info, err := os.Stat(appPath)

	if errors.Is(err, fs.ErrNotExist) || info.IsDir() == false {
		os.MkdirAll(appPath, 0755)
	}

	return appPath
}

func GetAppDBFile() string {
	dbPath := filepath.Join(GetAppFolder(), "remusika.db")

	return dbPath
}
