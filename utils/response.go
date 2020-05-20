package utils

import (
	"net/http"
	"log"
	"encoding/json"
)

func ToJson(writer http.ResponseWriter,data interface{}) {
	writer.Header().Set("Content-type","application/json")
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
