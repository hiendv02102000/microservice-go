package handler

import (
	"authentication/data/db"
	"authentication/data/dto"
	"authentication/handler/repository"
	"authentication/handler/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	Usecase usecase.UserUsecase
}

func NewHTTPHandler(db db.Database) *HTTPHandler {
	repo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(repo)
	return &HTTPHandler{Usecase: usecase}
}

func (h *HTTPHandler) Authenticate(c *gin.Context) {
	loginReq := dto.AuthenticateRequest{}
	err := c.ShouldBind(&loginReq)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data, err := h.Usecase.Authenticate(loginReq)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	c.JSON(http.StatusAccepted, data)
}
