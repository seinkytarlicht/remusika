package service

import (
	"github.com/seinkytarlicht/remusika/internal/model"
	"github.com/seinkytarlicht/remusika/internal/repository"
)

type MusicService interface {
	GetAll() []model.Music
	Get(uuid string) (model.Music, error)
	GetImage(uuid string) (string, error)
	ReloadFolder()
}

type MusicServiceImpl struct {
	Repo repository.MusicMemo
}

func (m *MusicServiceImpl) Get(uuid string) (model.Music, error) {
	data, err := m.Repo.Find(uuid)
	if err != nil {
		return model.Music{}, err
	}

	return data, nil
}

func (m *MusicServiceImpl) GetImage(uuid string) (string, error) {
	data, err := m.Repo.Find(uuid)
	if err != nil {
		return "", err
	}

	return data.TempImgPath, nil
}

func (m *MusicServiceImpl) GetAll() []model.Music {
	music := m.Repo.FindAll()

	return music
}

func (m *MusicServiceImpl) ReloadFolder() {
	m.Repo.ReloadFolder()
}

func NewMusicService(repository repository.MusicMemo) MusicService {
	return &MusicServiceImpl{
		Repo: repository,
	}
}
