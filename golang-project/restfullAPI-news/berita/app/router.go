package app

import (
	"restfullapi_news/berita/controller"
	"restfullapi_news/berita/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(beritaController controller.BeritaController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/news", beritaController.GetAll)
	router.GET("/news/category/:newsID", beritaController.GetByID)

	router.POST("/news", beritaController.Create)
	router.DELETE("/news/category/:newsID", beritaController.Delete)
	router.PUT("/news/category/:newsID", beritaController.Update)

	router.PanicHandler = exception.ErrorHandler

	return router
}
