package main

import (
	"context"
	"github.com/kataras/iris/v12"
	context2 "github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"go-practice/electricity-project/common"
	"go-practice/electricity-project/frontend/middlerware"
	"go-practice/electricity-project/frontend/web/controller"
	"go-practice/electricity-project/service"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./frontend/web/views", ".html").Layout("shared/layout.html").Reload(true))

	// 设置模版目录
	app.HandleDir("/public", iris.Dir("./frontend/web/public"))

	// 出现异常跳转到指定页面
	app.OnAnyErrorCode(func(context iris.Context) {
		context.ViewData("message", context.Values().GetStringDefault("message", "访问出错"))
		context.ViewLayout("")
		context.View("shared/error.html")
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		panic(err)
	}

	userService := service.NewUserService("user", db)
	userParty := app.Party("/user")
	user := mvc.New(userParty)
	user.Register(userService, ctx)
	user.Handle(new(controller.UserController))

	orderService := service.NewOrderService(db)

	// 注册控制权
	productServiceImp := service.NewProductService(db)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	productParty.Use(func(c *context2.Context) {
		middlerware.AuthProduct(c)
	})
	product.Register(ctx, productServiceImp, orderService)
	product.Handle(new(controller.ProductController))

	app.Run(iris.Addr("localhost:8082"))

}

//func main() {
//	app := iris.New()
//	app.Logger().SetLevel("debug")
//	template := iris.HTML("./frontend/web/views", ".html").
//		Layout("shared/layout.html").Reload(true)
//	app.RegisterView(template)
//	//设置模板目录
//	app.HandleDir("/assets", iris.Dir("./frontend/web/assets"))
//	//访问生成html静态文件
//	app.HandleDir("/html", iris.Dir("./frontend/web/generate/htmlProductOut"))
//	//出现异常跳转到指定页面
//	app.OnAnyErrorCode(func(ctx iris.Context) {
//		ctx.ViewData("message",
//			ctx.Values().GetStringDefault("message", "访问的页面出错！"))
//		ctx.ViewLayout("")
//		ctx.View("shared/error.html")
//	})
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	/*session := sessions.New(sessions.Config{
//		Cookie:  "helloword",
//		Expires: 60 * time.Minute,
//	})*/
//
//	db, err := common.NewMysqlConn()
//	if err != nil {
//		panic(err)
//	}
//
//	//sess := sessions.New(sessions.Config{
//	//	Cookie:  "helloword",
//	//	Expires: 60 * time.Minute,
//	//}
//
//	//注册控制权
//	userService := service.NewUserService("user", db)
//	userParty := app.Party("/user")
//	user := mvc.New(userParty)
//	user.Register(userService, ctx /*, session.Start*/)
//	user.Handle(new(controller.UserController))
//
//	orderService := service.NewOrderService(db)
//
//	productService := service.NewProductService(db)
//	productParty := app.Party("/product")
//	product := mvc.New(productParty)
//	//productParty.Use(func(c *context2.Context) {
//	//	middlerware.AuthProduct(c)
//	//})
//	//productParty.Use(middlerware.AuthProduct)
//	product.Register(ctx, orderService, productService /*, session.Start*/)
//	product.Handle(new(controller.ProductController))
//
//	//启动服务
//	app.Run(iris.Addr(":8082"))
//
//}
