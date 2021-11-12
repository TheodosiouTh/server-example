package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8080"

func main() {
	listeningPort := fmt.Sprintf(":%s", port)
	InitializeRoutes()
	log.Printf("Starting server on localhost:%s", port)
	err := http.ListenAndServe(listeningPort, nil)
	if err != nil {
		log.Fatalf("Error happened when starting server: %s\n", err.Error())
		return
	}
}
