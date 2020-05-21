package controllers
import (
	"net/http"
	"InjectMe/utils"
	"InjectMe/models"
	"io/ioutil"
	"encoding/json"
)

func GetUsers(writer http.ResponseWriter, request *http.Request){
	users, err := models.GetUsers()
	if err != nil {
		utils.ErrorResponse(writer, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(writer, users)
}

func PostUsers(writer http.ResponseWriter, req *http.Request){
	body,_ := ioutil.ReadAll(req.Body)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ErrorResponse(writer, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = models.NewUser(user)
	if err != nil {
		utils.ErrorResponse(writer, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(writer, utils.DefaultResponse{"User has been created", http.StatusCreated})
}
