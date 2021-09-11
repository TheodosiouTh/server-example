package message

import (
	"net/http"
)

func InitializeRoutes(){
	http.HandleFunc("/message/log", LogMessage)
}