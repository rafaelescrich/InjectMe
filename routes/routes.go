package routes
import (
	"net/http" 
	"encoding/json" 
	"github.com/gorilla/mux"
	"InjectMe/utils"
)

func NewRouter() *mux.Router {
	r:= mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", home).Methods("GET")
	return r
}

func home(writer http.ResponseWriter, r *http.Request){
	utils.ToJson(writer, struct{
		Message string `json:"message"`
	}{
		Message:"Inject Me! A SQLi vulnerable API",
	})
}
