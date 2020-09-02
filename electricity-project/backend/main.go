package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-practice/electricity-project/backend/web/controllers"
)

func main()  {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true))

	// 设置模版目录
	app.HandleDir("/assets", "./backend/web/assets")

	// 出现异常跳转到指定页面
	app.OnAnyErrorCode(func(context iris.Context) {
		context.ViewData("message", context.Values().GetStringDefault("message", "访问出错"))
		context.ViewLayout("")
		context.View("shared/error.html")
	})

	// 注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))

	app.Run(iris.Addr("localhost:8080"))

}
