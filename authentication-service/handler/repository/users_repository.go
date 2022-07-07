package repository

import (
	"authentication/data/db"
	"authentication/data/entity"
)

type UserRepository interface {
	FirstUser(condition entity.User) (entity.User, error)
}

type userRepository struct {
	DB db.Database
}

func NewUserRepository(db db.Database) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) FirstUser(condition entity.User) (entity.User, error) {
	user := entity.User{}
	err := u.DB.First(&user, condition)
	return user, err
}
