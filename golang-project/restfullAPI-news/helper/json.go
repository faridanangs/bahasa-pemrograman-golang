package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestToBody(request *http.Request, result interface{}) {
	decode := json.NewDecoder(request.Body)
	err := decode.Decode(&result)
	IfLogingErr(err, "error terjadi di decode create controller ReadRequestToBody")
}

func WriteRequestToBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("content-type", "application/json")
	encode := json.NewEncoder(writer)
	err := encode.Encode(response)
	IfLogingErr(err, "error terjadi di encode create controller WriteRequestToBody")
}
