package main

import (
	"lets_bid/routes"
	"log"
	"net/http"
)

func main() {

	// Handler
	http.Handle("/", routes.Handlers())

	log.Printf("Server up on port 8089")
	log.Fatal(http.ListenAndServe(":8089", nil))
}
