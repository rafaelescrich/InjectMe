package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OCFloripa/InjectMe/models"
	"github.com/OCFloripa/InjectMe/routes"
)

func main() {
	port := "3000"
	models.TestConnection()
	fmt.Printf("Inject Me running on port %s\n", port)
	routes := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), routes))
	fmt.Println("vim-go")
}
