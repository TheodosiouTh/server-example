package main

import (
	"fmt"
	"net/http"
)

const port = "8888";

func main() {
	listeningPort := fmt.Sprintf(":%s",port);
	err := http.ListenAndServe(listeningPort, nil);
	if err != nil {
		fmt.Printf("Error happened when starting server: %s", err.Error());
		return;
	}
}
