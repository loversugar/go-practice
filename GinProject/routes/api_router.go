package routes

import (
	"github.com/gin-gonic/gin"
	"go-practice/GinProject/controller"
)

func RegisterApiRouter(router *gin.Engine) {

	userRouter := router.Group("api")
	{
		userRouter.GET("user/getUserInfo", controller.GetUserInfo)
		userRouter.POST("user/createUser", controller.CreateUser)
	}

}
