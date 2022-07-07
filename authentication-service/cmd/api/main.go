package main

import (
	"authentication/handler"
	"log"
)

const webPort = ":80"

func main() {
	app := NewRouter()
	app.DB.Migrate()
	log.Printf("Starting authentication service on port %s\n", webPort)

	h := handler.NewHTTPHandler(app.DB)
	api := app.Engine.Group("")
	{
		api.POST("/authenticate", h.Authenticate)
	}
	// define http server
	app.Engine.Run(webPort)
}
