package rest

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/hadihammurabi/dummy-online-shop/app/delivery/rest/middleware"
	"github.com/hadihammurabi/dummy-online-shop/app/delivery/rest/request"
	"github.com/hadihammurabi/dummy-online-shop/app/delivery/rest/response"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"github.com/hadihammurabi/dummy-online-shop/app/service"
	"github.com/ngamux/ctx"
)

type OrderRest struct {
	mux     *ctx.CtxMux
	service *service.Service
}

func NewOrderRest(mux *ctx.CtxMux) *OrderRest {
	r := &OrderRest{mux: mux}
	r.route()
	return r
}

func (r *OrderRest) route() {
	r.mux.Get("/", r.Index)
	r.mux.Post("/", r.Store)
	r.mux.Get("/:id", r.Show)
}

func (r *OrderRest) getService() *service.Service {
	if r.service == nil {
		r.service = ioc.Use(service.Service{}).(*service.Service)
	}
	return r.service
}

func (r *OrderRest) Index(c *ctx.Context) error {
	authMiddleware := middleware.NewAuthMiddleware(r.service)
	user, err := authMiddleware.JWT(c)
	if err != nil {
		return response.Fail(c, http.StatusUnauthorized, err.Error())
	}

	orders, err := r.getService().Order.FindByUserID(context.Background(), user.ID)
	if err != nil {
		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	type rr struct {
		Product  string
		Price    uint
		Quantity uint
		Subtotal uint
	}
	orderResponse := make([]rr, 0)

	for _, order := range orders {
		orderResponse = append(orderResponse, rr{
			Product:  order.Product.Name,
			Price:    order.Product.Price,
			Quantity: order.Quantity,
			Subtotal: order.Quantity * order.Product.Price,
		})
	}

	return response.Success(c, http.StatusOK, orderResponse)
}

func (r *OrderRest) Show(c *ctx.Context) error {
	authMiddleware := middleware.NewAuthMiddleware(r.service)
	_, err := authMiddleware.JWT(c)
	if err != nil {
		return response.Fail(c, http.StatusUnauthorized, err.Error())
	}

	idFromUrl := c.GetParam("id")
	id, err := strconv.Atoi(idFromUrl)
	if err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	category, err := r.getService().Order.FindByID(context.Background(), uint(id))
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return response.Fail(c, http.StatusNotFound, err.Error())
		}

		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, category)
}

func (r *OrderRest) Store(c *ctx.Context) error {
	authMiddleware := middleware.NewAuthMiddleware(r.service)
	user, err := authMiddleware.JWT(c)
	if err != nil {
		return response.Fail(c, http.StatusUnauthorized, err.Error())
	}

	var ordersIn []*request.OrderCreate
	if err := c.GetJSON(&ordersIn); err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	orders := make([]*entity.Order, 0)
	for _, orderIn := range ordersIn {
		newOrder := orderIn.ToEntity()
		newOrder.User = user
		order, err := r.getService().Order.Create(context.Background(), newOrder)
		if err != nil {
			continue
		}

		orders = append(orders, order)
	}

	return response.Success(c, http.StatusOK, orders)
}
