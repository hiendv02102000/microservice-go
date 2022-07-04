package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

// 	out, _ := json.MarshalIndent(payload, "", "\t")
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusAccepted)
// 	w.Write(out)
// }
type HTTPHandler struct {
}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{}
}

func (h *HTTPHandler) Handle(c *gin.Context) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker1",
	}
	c.JSON(http.StatusAccepted, payload)
}
