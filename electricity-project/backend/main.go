package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-practice/electricity-project/backend/web/controllers"
	"go-practice/electricity-project/common"
	"go-practice/electricity-project/service"
)

func main()  {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true))

	// 设置模版目录
	app.HandleDir("/assets", iris.Dir("./backend/web/assets"))

	// 出现异常跳转到指定页面
	app.OnAnyErrorCode(func(context iris.Context) {
		context.ViewData("message", context.Values().GetStringDefault("message", "访问出错"))
		context.ViewLayout("")
		context.View("shared/error.html")
	})
	// 连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 注册控制权
	productServiceImp := service.NewProductService(db)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productServiceImp)
	product.Handle(new(controllers.ProductController))


	// 注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))

	app.Run(iris.Addr("localhost:8080"))

}
