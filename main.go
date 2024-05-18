package main

import (
	"log"
	"net/http"

	"github.com/ortizdavid/go-rest-concepts/config"
	"github.com/ortizdavid/go-rest-concepts/handlers"
)

func main() {
	mux := http.NewServeMux()
	
	handlers.RegisterRoutes(mux)
	log.Println("Listenning to: ", config.ListenAddr())
	http.ListenAndServe(config.ListenAddr(), mux)
}