package app

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)
func main() {
	app := iris.New()
	/**
	（可选）添加两个内置处理程序
		RECOVER：全局错误处理
		LOGGER：日志
	 */
	app.Use(recover.New())
	app.Use(logger.New())

	/**
	("/")给根路径注册路由。
		mvc.New(app).Handle(new(ExampleController))指在跟路由上注册控制器，访问路由就会进入对应的控制器。
		// http://localhost:8080
		// http://localhost:8080/ping
		// http://localhost:8080/hello
		// http://localhost:8080/custom_path
	 */
	mvc.New(app.Party("/")).Handle(new(ExampleController))
	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		panic(err)
	}
}
/**
	声明ExampleController控制器。并且服务于 "/", "/ping" and "/hello".
 */
type ExampleController struct{}
/**
	在控制器上注册方法路由
	比如注册get方法就需要以Get为开头，然后路由是其余部分。
 */

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}
// GetPing serves
// Method:   GET
// Resource: http://localhost:8080/ping
func (c *ExampleController) GetPing() string {
	return "pong"
}
// GetHello serves
// Method:   GET
// Resource: http://localhost:8080/hello
func (c *ExampleController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}
// Get GetMovieName serves
// Method:   GET
// Resource: http://localhost:8080/movie/name
func (c *ExampleController) GetMovieName(username string) string {
	return username
}
// GetUserBy serves 动态路由
// Method:   GET
// Resource: http://localhost:8080/user/:string
// 在控制器中使用Get{router name}By，方法的参数就是你传入的子路径
func (c *ExampleController) GetUserBy(username string) string {
	return username
}

//在控制器适应主应用程序之前调用一次BeforeActivation
//当然在服务器运行之前。
//在版本9之后，您还可以为特定控制器的方法添加自定义路由。
//在这里您可以注册自定义方法的处理程序
//使用带有`ca.Router`的标准路由器做一些你可以做的事情，没有mvc，
//并添加将绑定到控制器的字段或方法函数的输入参数的依赖项。
func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		ctx.Next()
	}
	b.Handle("GET", "/custom_path", "CustomHandlerWithoutFollowingTheNamingGuide", anyMiddlewareHere)

	//甚至添加基于此控制器路由器的全局中间件，
	//在这个例子中是根“/”：
	// b.Router（）。使用（myMiddleware）
}
// CustomHandlerWithoutFollowingTheNamingGuide serves
// Method:   GET
// Resource: http://localhost:8080/custom_path
func (c *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	return "hello from the custom handler without following the naming guide"
}


/**
	其他方法：
	func (c *ExampleController) Post() {}
	func (c *ExampleController) Put() {}
	func (c *ExampleController) Delete() {}
	func (c *ExampleController) Connect() {}
	func (c *ExampleController) Head() {}
	func (c *ExampleController) Patch() {}
	func (c *ExampleController) Options() {}
	func (c *ExampleController) Trace() {}
*/
/*
   func (c *ExampleController) All() {}
   //        OR
   func (c *ExampleController) Any() {}

   func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
       // 1 -> the HTTP Method
       // 2 -> the route's path
       // 3 -> this controller's method name that should be handler for that route.
       b.Handle("GET", "/mypath/{param}", "DoIt", optionalMiddlewareHere...)
   }
   // After activation, all dependencies are set-ed - so read only access on them
   // but still possible to add custom controller or simple standard handlers.
   func (c *ExampleController) AfterActivation(a mvc.AfterActivation) {}
*/


