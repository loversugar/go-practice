package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-practice/electricity-project/service"
)

type OrderController struct {
	Ctx          iris.Context
	OrderService service.IOrderService
}

func (o *OrderController) Get() mvc.View {
	orderArray, err := o.OrderService.GetAllOrder()

	if err != nil {
		o.Ctx.Application().Logger().Debug("查询订单详情失败")
	}

	return mvc.View{
		Name: "order/view.html",
		Data: iris.Map{
			"order": orderArray,
		},
	}
}
