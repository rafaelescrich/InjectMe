package controllers
import (
	"net/http"
	"InjectMe/utils"
	"InjectMe/models"
)

func GetUsers(writer http.ResponseWriter, request *http.Request){
	users, err := models.GetUsers()
	if err != nil {
		utils.ErrorResponse(writer, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(writer, users)
}
