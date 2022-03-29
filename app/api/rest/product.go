package rest

import (
	"context"
	"net/http"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/service"
	"github.com/ngamux/ctx"
	"github.com/ngamux/ngamux"
)

type ProductRest struct {
	mux     *ctx.CtxMux
	service *service.Service
}

func NewProductRest(mux *ctx.CtxMux) *ProductRest {
	r := &ProductRest{mux: mux}
	r.route()
	return r
}

func (r *ProductRest) route() {
	r.mux.Get("/", r.Index)
}

func (r *ProductRest) getService() *service.Service {
	if r.service == nil {
		r.service = ioc.Use(service.Service{}).(*service.Service)
	}
	return r.service
}

func (r *ProductRest) Index(c *ctx.Context) error {
	products, err := r.getService().Product.All(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ngamux.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ngamux.Map{
		"products": products,
	})
}
