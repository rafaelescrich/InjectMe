package main

import ("fmt" 
	"log" 
	"InjectMe/routes" 
	"InjectMe/models"
	"net/http"
)

func main() {
	port :="3000"
	models.TestConnection()
	fmt.Printf("Inject Me running on port %s\n", port)
	routes := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port),routes))
	fmt.Println("vim-go")
}
