package rest

import (
	"context"
	"net/http"

	"github.com/hadihammurabi/dummy-online-shop/app/delivery/rest/request"
	"github.com/hadihammurabi/dummy-online-shop/app/delivery/rest/response"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/service"
	"github.com/ngamux/ctx"
)

type AuthRest struct {
	mux     *ctx.CtxMux
	service *service.Service
}

func NewAuthRest(mux *ctx.CtxMux) *AuthRest {
	r := &AuthRest{mux: mux}
	r.route()
	return r
}

func (r *AuthRest) route() {
	r.mux.Post("/register", r.Register)
	r.mux.Post("/login", r.Login)
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
	return nil
}
