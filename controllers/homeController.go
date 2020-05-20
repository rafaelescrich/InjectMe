package controllers

import (
	"net/http"
	"InjectMe/utils"
)
	

func GetHome(writer http.ResponseWriter, r *http.Request) {
	utils.ToJson(writer, struct{
		Message string `json:"message"`
	}{
	Message:"Inject Me - A SQLi vulnerable API",
	})
}
