package rest

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/hadihammurabi/dummy-online-shop/app/api/rest/response"
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
	r.mux.Get("/:id", r.Show)
}

func (r *CategoryRest) getService() *service.Service {
	if r.service == nil {
		r.service = ioc.Use(service.Service{}).(*service.Service)
	}
	return r.service
}

func (r *CategoryRest) Index(c *ctx.Context) error {
	products, err := r.getService().Category.All(context.Background())
	if err != nil {
		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, products)
}

func (r *CategoryRest) Show(c *ctx.Context) error {
	idFromUrl := c.GetParam("id")
	id, err := strconv.Atoi(idFromUrl)
	if err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	product, err := r.getService().Category.FindByID(context.Background(), uint(id))
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return response.Fail(c, http.StatusNotFound, err.Error())
		}

		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, product)
}
