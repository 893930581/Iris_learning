package services

import (
	"fmt"
	"learn_Iris/respositories"
)

type MovieService interface {
	//业务逻辑类
	ShowMovieName() string
}

type MovieServiceManger struct {
	repo respositories.MovieRespository
}

func NewMovieServiceManger(repo respositories.MovieRespository) MovieService {
	return &MovieServiceManger{
		repo:repo,
	}
}

func (m *MovieServiceManger) ShowMovieName() string {
	fmt.Println("我们获取的视频" + m.repo.GetMovieName())
	return "我们获取的视频" + m.repo.GetMovieName()
}

