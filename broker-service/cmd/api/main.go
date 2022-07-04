package main

import (
	"broker/cmd/router"
	"broker/handler"
	"log"
)

const webPort = ":80"

func main() {
	app := router.NewRouter()

	log.Printf("Starting broker service on port %s\n", webPort)
	h := handler.NewHTTPHandler()
	api := app.Engine.Group("")
	{
		api.POST("/", h.Handle)
	}
	// define http server
	app.Engine.Run(webPort)
}
