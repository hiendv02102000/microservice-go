package main

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}
type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewRouter() Router {
	return Router{
		Engine: gin.Default(),
	}
}
