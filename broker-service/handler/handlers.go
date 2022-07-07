package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
}

func NewHTTPHandler() *HTTPHandler {
	return &HTTPHandler{}
}

func (h *HTTPHandler) Broker(c *gin.Context) {
	payload := BaseResponse{
		Status: http.StatusAccepted,
		Error:  nil,
		Result: "Hit the broker",
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusAccepted, payload)
}
func (h *HTTPHandler) HandleSubmission(c *gin.Context) {
	requestPayload := RequestPayload{}
	err := c.ShouldBind(&requestPayload)
	if err != nil {
		data := BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	switch requestPayload.Action {
	case "auth":
		{
			h.authenticate(c, requestPayload.Auth)
			break
		}

	default:
		{
			data := BaseResponse{
				Status: http.StatusBadRequest,
				Error:  errors.New("unknown action"),
			}
			c.JSON(http.StatusBadRequest, data)
		}

	}
}
func (h *HTTPHandler) authenticate(c *gin.Context, a AuthPayload) {
	jsonData, _ := json.MarshalIndent(a, "", "\t")
	response, err := http.Post("http://authentication-service/authenticate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		data := BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := BaseResponse{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		data := BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	c.JSON(http.StatusOK, data)
}
