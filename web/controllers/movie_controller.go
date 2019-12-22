package controllers

import (
	"github.com/kataras/iris/mvc"
	"learn_Iris/respositories"
	"learn_Iris/services"
)

type MovieController struct {

}

func (c *MovieController) GetMovieName() mvc.View {
	movieRepository := respositories.NewMovieManager()
	//创建数据库操作对象
	movieService := services.NewMovieServiceManger(movieRepository)
	//创建业务逻辑对象
	name := movieService.ShowMovieName()
	//调用业务逻辑代码获取数据
	return mvc.View{
		Name:"movie/index.html",
		Data:name,
	}
}