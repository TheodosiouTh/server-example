package image

import (
	"image/png"
	"log"
	"net/http"
	"strconv"
)

func BlurImageHandler(response http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(5 * 1024 * 1024)
	if err != nil {
		log.Print("Could not parse request form data")
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	imageData, _, err := request.FormFile("image")
	if err != nil {
		log.Print("Could not parse image from request form data")
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	standardDeviation, err := strconv.ParseFloat(request.FormValue("standardDeviation"), 64)
	if err != nil {
		log.Print("Could not parse standardDeviation from request form data, setting default value of 3.0")
		standardDeviation = 3.0
	}

	err = imageData.Close()
	if err != nil {
		log.Print("Could not parse image from request data")
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	image, err := png.Decode(imageData)
	if err != nil {
		log.Print("Could not decode image from request data")
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	imageChannel := make(chan []byte)
	go BlurImage(image, standardDeviation, imageChannel)
	bluredImageBurffer := <-imageChannel

	response.Header().Set("Content-Type", "image/png")
	_, err = response.Write(bluredImageBurffer)
	if err != nil {
		log.Print("Could not send blured image to client")
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
