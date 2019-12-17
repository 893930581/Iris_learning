package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	//日志输出等级
	app.RegisterView(iris.HTML("./web/views",".html"))
	//注册视图层
	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		panic(fmt.Sprintf("start server failed, err:%v",err))
	}
	//监听端口
}