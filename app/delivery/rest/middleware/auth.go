package middleware

import (
	"errors"
	"strings"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"github.com/hadihammurabi/dummy-online-shop/app/service"
	"github.com/ngamux/ctx"
)

type AuthMiddleware struct {
	service *service.Service
}

func NewAuthMiddleware(service *service.Service) *AuthMiddleware {
	return &AuthMiddleware{
		service: service,
	}
}

func (m *AuthMiddleware) getService() *service.Service {
	if m.service == nil {
		m.service = ioc.Use(service.Service{}).(*service.Service)
	}
	return m.service
}

func (m *AuthMiddleware) JWT(c *ctx.Context) (*entity.User, error) {
	authHeader := c.Request().Header.Get("authorization")
	if authHeader == "" {
		return nil, errors.New("unauthorized")
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) <= 1 || bearerToken[0] != "Bearer" {
		return nil, errors.New("invalid token")
	}

	return m.getService().Auth.Info(c.Context(), bearerToken[1])
}
