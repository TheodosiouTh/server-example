package main

import (
	"server-example/api/calculator"
	"server-example/api/image"
	"server-example/api/message"
)

func InitializeRoutes() {
	image.InitializeRoutes()
	message.InitializeRoutes()
	calculator.InitializeRoutes()
}
