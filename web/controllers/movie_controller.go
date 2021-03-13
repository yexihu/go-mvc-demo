package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"go-mvc-demo/repositories"
	"go-mvc-demo/services"
)

type MovieController struct {

}

func (c *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMoieServiceManager(movieRepository)

	movieResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: movieResult,
	}
}