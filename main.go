package main

import (
	"lets_bid/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {

	// Handler
	http.Handle("/", routes.Handlers())

	log.Printf("Server up on port 8089")
	log.Fatal(http.ListenAndServe(":8089", handlers.CORS()(r)))
}
