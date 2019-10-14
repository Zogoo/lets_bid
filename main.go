package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading env file")
	}

	connectDb()

}
