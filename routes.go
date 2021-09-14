package main

import (
	"github.com/TheodosiouTh/server-example/api/calculator"
	"github.com/TheodosiouTh/server-example/api/image"
	"github.com/TheodosiouTh/server-example/api/message"
)

func InitializeRoutes() {
	image.InitializeRoutes()
	message.InitializeRoutes()
	calculator.InitializeRoutes()
}
