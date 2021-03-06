package rest

import (
	"context"
	"net/http"

	"github.com/hadihammurabi/dummy-online-shop/app/delivery/rest/middleware"
	"github.com/hadihammurabi/dummy-online-shop/app/delivery/rest/request"
	"github.com/hadihammurabi/dummy-online-shop/app/delivery/rest/response"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/service"
	"github.com/ngamux/ctx"
)

type AuthRest struct {
	mux            *ctx.CtxMux
	service        *service.Service
	authMiddleware *middleware.AuthMiddleware
}

func NewAuthRest(mux *ctx.CtxMux) *AuthRest {
	r := &AuthRest{mux: mux, authMiddleware: &middleware.AuthMiddleware{}}
	r.route()
	return r
}

func (r *AuthRest) route() {
	r.mux.Post("/register", r.Register)
	r.mux.Post("/login", r.Login)
	r.mux.Get("/info", r.Info)
}

func (r *AuthRest) getService() *service.Service {
	if r.service == nil {
		r.service = ioc.Use(service.Service{}).(*service.Service)
	}
	return r.service
}

func (r *AuthRest) Register(c *ctx.Context) error {
	var reqBody request.AuthRegisterRequest
	if err := c.GetJSON(&reqBody); err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	if err := reqBody.Validate(); err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	user, err := r.getService().User.Create(context.Background(), reqBody.ToEntity())
	if err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	return response.Success(c, http.StatusOK, user)
}

func (r *AuthRest) Login(c *ctx.Context) error {
	var reqBody request.AuthLoginRequest
	if err := c.GetJSON(&reqBody); err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	if err := reqBody.Validate(); err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	token, err := r.getService().Auth.Login(context.Background(), reqBody.Email, reqBody.Password)
	if err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	return response.Success(c, http.StatusOK, token)
}

func (r *AuthRest) Info(c *ctx.Context) error {
	authMiddleware := middleware.NewAuthMiddleware(r.getService())
	user, err := authMiddleware.JWT(c)
	if err != nil {
		return response.Fail(c, http.StatusUnauthorized, err.Error())
	}
	return response.Success(c, http.StatusOK, user)
}
