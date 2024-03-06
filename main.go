package main

import (
	"log"
	"net/http"
	"github.com/ortizdavid/go-rest-concepts/handlers"
)

func main() {
	mux := http.NewServeMux()

	handlers.RegisterRoutes(mux)

	log.Println("Listenning to: http://127.0.0.1:8000")

	http.ListenAndServe(":8000", mux)
}