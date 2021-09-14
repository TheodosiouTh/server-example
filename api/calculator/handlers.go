package calculator

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type sumBody struct {
	Numbers []int
}

func DoubleHandler(response http.ResponseWriter, request *http.Request) {
	var body sumBody
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		log.Print("Could not decode request body")
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	result := Sum(body.Numbers)
	fmt.Fprintln(response, result)
}
