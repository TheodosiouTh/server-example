package calculator

import (
	"net/http"
)

func InitializeRoutes() {
	http.HandleFunc("/calculator/sum-integers", DoubleHandler)
}
