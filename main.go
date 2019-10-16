package main

import (
	"lets_bid/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {

	routers := routes.Handlers()
	// Handler
	http.Handle("/", routers)

	log.Printf("Server up on port 8089")
	log.Fatal(http.ListenAndServe(":8089", handlers.CORS()(routers)))
}
