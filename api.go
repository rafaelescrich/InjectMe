package main

import ("fmt" 
	"log" 
	"InjectMe/routes" 
	"net/http"
)

func main() {
	port :="3000"
	fmt.Printf("Inject Me running on port %s\n", port)
	routes := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port),routes))
	fmt.Println("vim-go")
}
