package routes
import (
	"github.com/gorilla/mux"
	"InjectMe/controllers"
)

func NewRouter() *mux.Router {
	r:= mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.GetHome).Methods("GET")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users", controllers.PostUsers).Methods("POST")
	return r
}
