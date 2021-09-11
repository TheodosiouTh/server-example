package message

import (
	"encoding/json"
	"log"
	"net/http"
)

type logMessageBody struct {
	Message string
}

func LogMessage(response http.ResponseWriter, request *http.Request) {
	var  body  logMessageBody
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		log.Print("Could not decode request body")
        http.Error(response, err.Error(), http.StatusBadRequest);
		return
	}

	log.Print(body.Message);
	response.Write([]byte("Message was logged successfully!"));
}