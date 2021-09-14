package calculator

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func SumIntegerHandler(response http.ResponseWriter, request *http.Request) {
	sentNumbers := request.URL.Query()["numbers"]

	var numbers []int
	for _, numberString := range sentNumbers {
		number, err := strconv.ParseInt(numberString, 10, 64)
		if err != nil {
			log.Print("Could not decode request body")
			http.Error(response, err.Error(), http.StatusBadRequest)
			fmt.Fprintln(response, "No numbers or special characters allowed.")
			return
		}
		numbers = append(numbers, int(number))
	}

	result := Sum(numbers)
	fmt.Fprintln(response, result)
}
