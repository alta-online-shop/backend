package service

import (
	"context"
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type AuthService interface {
	Login(c context.Context, email string, password string) (string, error)
}

type authService struct {
	repo *repository.Repository
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) getRepo() *repository.Repository {
	if s.repo == nil {
		s.repo = ioc.Use(repository.Repository{}).(*repository.Repository)
	}

	return s.repo
}

func (s *authService) Login(c context.Context, email string, password string) (string, error) {
	user, err := s.getRepo().User.FindByEmail(c, email)
	if err != nil {
		return "", err
	}

	if password != user.Password {
		return "", errors.New("email or password is invalid")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.JWTClaims{
		Fullname: user.Fullname,
		Email:    user.Email,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("APP_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
