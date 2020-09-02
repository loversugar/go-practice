package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"go-practice/electricity-project/repositories"
	"go-practice/electricity-project/service"
)

type MovieController struct {

}

func (c *MovieController) Get() mvc.View  {
	movieRep := repositories.NewMovieManager()
	moviceService := service.NewMovieServiceManager(movieRep)
	result := moviceService.ShowMovieName()

	return mvc.View{
		Name: "movie/index.html",
		Data: result,
	}
}
