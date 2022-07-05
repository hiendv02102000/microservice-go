package main

import (
	"authentication/data/db"
	"authentication/handler"
	"log"
)

const webPort = ":80"

func main() {
	app := NewRouter()

	log.Printf("Starting authentication service on port %s\n", webPort)
	_, err := db.NewDB()
	log.Println(err)
	h := handler.NewHTTPHandler()
	api := app.Engine.Group("")
	{
		api.POST("/", h.Broker)
	}
	// define http server
	app.Engine.Run(webPort)
}
