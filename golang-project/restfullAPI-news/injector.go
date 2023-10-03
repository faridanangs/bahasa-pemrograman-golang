//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"restfullapi_news/berita/app"
	"restfullapi_news/berita/controller"
	"restfullapi_news/berita/repository"
	"restfullapi_news/berita/service"
	"restfullapi_news/middleware"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewBeritaRepository,
	wire.Bind(new(repository.BeritaRepository), new(repository.BeritaRepositoryIplm)),
	service.NewBeritaService,
	wire.Bind(new(service.BeritaService), new(service.BeritaServiceIplm)),
	controller.NewBeritaController,
	wire.Bind(new(controller.BeritaController), new(controller.BeritaControllerIplm)),
)

func WireInjectionBerita() *http.Server {
	wire.Build(
		categorySet,
		validator.New,
		app.ConnectToDB,
		app.NewRouter,
		app.NewServerNangs,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
	)
	return nil
}
