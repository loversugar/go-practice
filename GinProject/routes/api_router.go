package routes

import (
	"github.com/gin-gonic/gin"
	"go-practice/GinProject/controller"
)

func RegisterApiRouter(router *gin.Engine) {

	apiRouter := router.Group("api")
	{
		apiRouter.GET("user/getUserInfo", controller.GetUserInfo)
	}
}
