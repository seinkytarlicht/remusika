package repository

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/seinkytarlicht/remusika/helper"
	"github.com/seinkytarlicht/remusika/internal/model"
	"go.senan.xyz/taglib"
)

type MusicMemo interface {
	FindAll() []model.Music
	Find(uuid string) (model.Music, error)
	ReloadFolder()
}

type MusicMemoImpl struct {
	Musics   []model.Music
	Playlist []string
}

func (m *MusicMemoImpl) FindAll() []model.Music {
	return m.Musics
}

func (m *MusicMemoImpl) Find(uuid string) (model.Music, error) {
	index := slices.IndexFunc(m.Musics, func(mu model.Music) bool {
		return mu.Uuid == uuid
	})

	if index != -1 {
		return m.Musics[index], nil
	} else {
		return model.Music{}, errors.New("Music not Found")
	}
}

func (m *MusicMemoImpl) ReloadFolder() {
	helper.CleanUpTemp()

	var music []model.Music
	var listpath []string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to detect home directory: %v", err)
	}
	musicFolder := filepath.Join(homeDir, "Music")

	errW := filepath.WalkDir(musicFolder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))

		switch ext {
		case ".mp3", ".flac", ".wav", ".ogg", ".opus", ".m4a":
			listpath = append(listpath, path)
		}

		return nil
	})
	if errW != nil {
		log.Fatal(errW)
	}

	for _, lp := range listpath {
		hashBytes := sha1.Sum([]byte(lp))
		uuid := hex.EncodeToString(hashBytes[:])[:12]

		tags, err := taglib.ReadTags(lp)
		if err != nil {
			log.Fatal(err)
		}

		musicName := strings.Join(tags[taglib.Title], ", ")
		musicArtist := strings.Join(tags[taglib.Artist], ", ")
		musicAlbum := strings.Join(tags[taglib.Album], ", ")
		playlist := "No Playlist"
		tmpImg, _ := _CreateImageTemp(lp)
		imageUrl := ""
		audioUrl := "/api/music/stream/" + uuid

		if musicName == "" {
			musicName = filepath.Base(lp)
		}

		if musicArtist == "" {
			musicArtist = "No Artist"
		}

		if musicAlbum == "" {
			musicAlbum = "No Album"
		}

		// Playlist
		relPath, errR := filepath.Rel(musicFolder, lp)
		if errR != nil {
			log.Fatal(errR)
		}
		parts := strings.Split(relPath, string(os.PathSeparator))
		if len(parts) > 1 {
			playlist = parts[0]
			m.Playlist = append(m.Playlist, playlist)
		}

		if tmpImg != "" {
			imageUrl = "/api/music/image/" + uuid
		}

		music = append(music, model.Music{
			Uuid:        uuid,
			Title:       musicName,
			Artist:      musicArtist,
			Playlist:    playlist,
			Album:       musicAlbum,
			AudioUrl:    audioUrl,
			ImageUrl:    imageUrl,
			TempImgPath: tmpImg,
			Path:        lp,
		})
	}

	m.Musics = music
}

func _CreateImageTemp(lp string) (string, error) {
	imageBytes, err := taglib.ReadImage(lp)
	if err != nil {
		return "", err
	}

	filename, err := helper.CreateImageTemp(imageBytes)

	if err != nil {
		return "", err
	}

	return filename, nil
}

func NewMusicMemo() MusicMemo {
	var musicMemoImpl MusicMemo = &MusicMemoImpl{}
	musicMemoImpl.ReloadFolder()

	return musicMemoImpl
}
