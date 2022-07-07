package usecase

import (
	"authentication/data/dto"
	"authentication/data/entity"
	"authentication/handler/repository"
	"errors"
	"net/http"
)

type UserUsecase interface {
	Authenticate(reqLogin dto.AuthenticateRequest) (dto.BaseResponse, error)
}

type userUsercase struct {
	repo repository.UserRepository
}

func (u *userUsercase) Authenticate(reqLogin dto.AuthenticateRequest) (dto.BaseResponse, error) {
	user, err := u.repo.FirstUser(entity.User{
		Email: reqLogin.Email,
	})
	if err != nil {
		return dto.BaseResponse{}, err
	}
	if user.ID == 0 {
		return dto.BaseResponse{}, errors.New("record not found")
	}
	if user.Password != reqLogin.Password {
		return dto.BaseResponse{}, errors.New("not match password")
	}
	return dto.BaseResponse{
		Status: http.StatusOK,
		Result: "Authen",
	}, nil
}
func NewUserUsecase(r repository.UserRepository) *userUsercase {
	return &userUsercase{
		repo: r,
	}
}
