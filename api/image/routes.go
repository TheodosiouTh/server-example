package image

import (
	"net/http"
)

func InitializeRoutes() {
	http.HandleFunc("/image/blur", BlurImageHandler)
}
