package main

import (
	"authentication/data/db"
	"log"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	DB     db.Database
}

func NewRouter() Router {
	database, err := db.NewDB()
	if err != nil {
		log.Println(err)
		return Router{}
	}
	return Router{
		Engine: gin.Default(),
		DB:     database,
	}
}
