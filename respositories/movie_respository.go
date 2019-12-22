package respositories

import (
	"learn_Iris/datamodels"
	"learn_Iris/datasource"
)


//业务层
type MovieRespository interface {
	GetMovieName() string
}

type MovieRespositoryManager struct {}

func NewMovieManager() MovieRespository {
	//数据库操作类
	return &MovieRespositoryManager{}
}

func (m *MovieRespositoryManager) GetMovieName() string {
	movie := &datamodels.Movie{Name:datasource.Movies[1].Name}
	//模拟数据库操作
	return movie.Name
}