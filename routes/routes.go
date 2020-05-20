package routes
import ("net/http" "encoding/json" "github.com/gorilla/mux")

func NewRouter() *mux.Router {
	r:= mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", home).Methods("GET")
	return r
}

func home(w http.ResponseWriter, r *http.Requests){
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(struct{
		Message string `json:"message"`
	}{
		Message: "Hello World!",
	})
}
