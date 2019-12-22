package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"learn_Iris/web/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	//日志输出等级
	app.RegisterView(iris.HTML("./web/views",".html"))
	//注册视图层
	mvc.New(app).Handle(new(controllers.MovieController))
	//注册控制器
	err := app.Run(iris.Addr(":8888"))
	//监听端口
	if err != nil {
		panic(fmt.Sprintf("start server failed, err:%v",err))
	}
}