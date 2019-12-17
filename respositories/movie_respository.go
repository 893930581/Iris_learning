package respository

import "github.com/kataras/iris/_examples/mvc/overview/datamodels"
//业务层
type MovieRespository interface {
	GetMovieName() string
}

type MovieManager struct {

}

func NewMovieManager() MovieRespository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	movie := &datamodels.Movie{Name:"流星花园"}
	return movie.Name
}