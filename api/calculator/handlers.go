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
		number, err := strconv.Atoi(numberString)
		if err != nil {
			log.Print("Could not decode request body")
			http.Error(response, "No numbers or special characters allowed.", http.StatusBadRequest)
			return
		}
		numbers = append(numbers, int(number))
	}

	result := Sum(numbers)
	fmt.Fprint(response, result)
}
