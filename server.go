package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8888";

func main() {
	listeningPort := fmt.Sprintf(":%s",port);
	InitializeRoutes();
	err := http.ListenAndServe(listeningPort, nil);
	if err != nil {
		log.Fatalf("Error happened when starting server: %s\n", err.Error());
		return;
	}
}
