package utils

import (
	"net/http"
	"log"
	"encoding/json"
)

func ErrorResponse(writer http.ResponseWriter, err error, status int){
	writer.WriteHeader(status)
	ToJson(writer, struct{
		Message string `json:"message"`
	}{
		Message: err.Error(),
	})
}

func ToJson(writer http.ResponseWriter,data interface{}) {
	writer.Header().Set("Content-type","application/json")
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
