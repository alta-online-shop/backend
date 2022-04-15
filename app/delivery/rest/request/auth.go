package request

import (
	"errors"

	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type AuthRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

func (r AuthRegisterRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}

	if r.Password == "" {
		return errors.New("password is required")
	}

	if r.Fullname == "" {
		return errors.New("fullname is required")
	}

	return nil
}

func (r AuthRegisterRequest) ToEntity() *entity.User {
	return &entity.User{
		Fullname: r.Fullname,
		Email:    r.Email,
		Password: r.Password,
	}
}

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r AuthLoginRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}

	if r.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
