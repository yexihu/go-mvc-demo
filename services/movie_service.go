package services

import (
	"go-mvc-demo/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repositories.MovieReposity
}

func NewMoieServiceManager(repo repositories.MovieReposity) MovieService{
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string {
	text := "获取视频名称："+ m.repo.GetMovieName()
	return text
}