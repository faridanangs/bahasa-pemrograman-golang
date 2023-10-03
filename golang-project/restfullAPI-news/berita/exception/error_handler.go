package exception

import (
	"net/http"
	"restfullapi_news/berita/model/web"
	"restfullapi_news/helper"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if validationError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)

}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Status: "BAD REQUEST",
			Code:   http.StatusBadRequest,
			Data:   exception.Error(),
		}
		helper.WriteRequestToBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Status: "Internal error",
		Code:   http.StatusInternalServerError,
		Data:   err,
	}
	helper.WriteRequestToBody(writer, webResponse)
}
