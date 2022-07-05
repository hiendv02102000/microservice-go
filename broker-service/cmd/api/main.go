package main

import (
	"broker/handler"
	"log"
)

const webPort = ":80"

func main() {
	app := NewRouter()

	log.Printf("Starting broker service on port %s\n", webPort)
	h := handler.NewHTTPHandler()
	api := app.Engine.Group("")
	{
		api.POST("/", h.Broker)
	}
	// define http server
	app.Engine.Run(webPort)
}
