package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestToBody(request *http.Request, result interface{}) {
	// ambil nilai dari request di dalam body
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteRequestToBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
