package rest

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/hadihammurabi/dummy-online-shop/app/delivery/rest/request"
	"github.com/hadihammurabi/dummy-online-shop/app/delivery/rest/response"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/service"
	"github.com/ngamux/ctx"
)

type CategoryRest struct {
	mux     *ctx.CtxMux
	service *service.Service
}

func NewCategoryRest(mux *ctx.CtxMux) *CategoryRest {
	r := &CategoryRest{mux: mux}
	r.route()
	return r
}

func (r *CategoryRest) route() {
	r.mux.Get("/", r.Index)
	r.mux.Post("/", r.Store)
	r.mux.Get("/:id", r.Show)
}

func (r *CategoryRest) getService() *service.Service {
	if r.service == nil {
		r.service = ioc.Use(service.Service{}).(*service.Service)
	}
	return r.service
}

func (r *CategoryRest) Index(c *ctx.Context) error {
	categories, err := r.getService().Category.All(context.Background())
	if err != nil {
		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, categories)
}

func (r *CategoryRest) Show(c *ctx.Context) error {
	idFromUrl := c.GetParam("id")
	id, err := strconv.Atoi(idFromUrl)
	if err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	category, err := r.getService().Category.FindByID(context.Background(), uint(id))
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return response.Fail(c, http.StatusNotFound, err.Error())
		}

		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, category)
}

func (r *CategoryRest) Store(c *ctx.Context) error {
	var categoryIn *request.CategoryCreate
	if err := c.GetJSON(&categoryIn); err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	category, err := r.getService().Category.Create(context.Background(), categoryIn.ToEntity())
	if err != nil {
		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, category)
}
