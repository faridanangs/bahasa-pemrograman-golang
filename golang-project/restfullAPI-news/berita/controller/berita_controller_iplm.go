package controller

import (
	"net/http"
	"restfullapi_news/berita/model/web"
	"restfullapi_news/berita/service"
	"restfullapi_news/helper"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BeritaControllerIplm struct {
	BeritaService service.BeritaService
}

func NewBeritaController(beritaService service.BeritaService) *BeritaControllerIplm {
	return &BeritaControllerIplm{
		BeritaService: beritaService,
	}
}

func (controller *BeritaControllerIplm) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	beritaCreateRequest := web.BeritaCreateRequest{}
	helper.ReadRequestToBody(request, &beritaCreateRequest)

	response := controller.BeritaService.Create(request.Context(), beritaCreateRequest)

	webResponse := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   response,
	}
	helper.WriteRequestToBody(writer, webResponse)

}
func (controller *BeritaControllerIplm) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	beritaUpdateRequest := web.BeritaUpdateRequest{}
	helper.ReadRequestToBody(request, &beritaUpdateRequest)

	paramsID := params.ByName("newsID")
	id, err := strconv.Atoi(paramsID)
	helper.IfLogingErr(err, "error terjadi di strconv.Atoi update controller")

	beritaUpdateRequest.ID = id

	response := controller.BeritaService.Update(request.Context(), beritaUpdateRequest)

	webResponse := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   response,
	}

	helper.WriteRequestToBody(writer, webResponse)
}
func (controller *BeritaControllerIplm) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	paramsID := params.ByName("newsID")
	id, err := strconv.Atoi(paramsID)
	helper.IfLogingErr(err, "error terjadi di strconv.Atoi delete controller")

	controller.BeritaService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
	}
	helper.WriteRequestToBody(writer, webResponse)
}
func (controller *BeritaControllerIplm) GetByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paramsID := params.ByName("newsID")
	id, err := strconv.Atoi(paramsID)
	helper.IfLogingErr(err, "error terjadi di strconv.Atoi delete controller")

	beritaResponse := controller.BeritaService.GetByID(request.Context(), id)
	webResponse := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   beritaResponse,
	}
	helper.WriteRequestToBody(writer, webResponse)
}
func (controller *BeritaControllerIplm) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	beritaResponse := controller.BeritaService.GetAll(request.Context())
	webResponse := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   beritaResponse,
	}
	helper.WriteRequestToBody(writer, webResponse)
}
