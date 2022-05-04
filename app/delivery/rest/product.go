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
	"github.com/hadihammurabi/dummy-online-shop/app/service"
	"github.com/ngamux/ctx"
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
	r.mux.Post("/", r.Store)
	r.mux.Get("/:id", r.Show)
	r.mux.Delete("/:id", r.Destroy)

	r.mux.Get("/:id/ratings", r.GetRatings)
	r.mux.Post("/:id/ratings", r.SetRatings)
}

func (r *ProductRest) getService() *service.Service {
	if r.service == nil {
		r.service = ioc.Use(service.Service{}).(*service.Service)
	}
	return r.service
}

func (r *ProductRest) Index(c *ctx.Context) error {
	categoryFromQuery := c.GetQuery("category")
	if categoryFromQuery == "" {
		categoryFromQuery = c.GetQuery("c")
	}

	categoryID, err := strconv.Atoi(categoryFromQuery)
	if err != nil {
		categoryID = 0
	}

	if categoryID == 0 {
		products, err := r.getService().Product.All(context.Background())
		if err != nil {
			return response.Fail(c, http.StatusInternalServerError, err.Error())
		}

		return response.Success(c, http.StatusOK, products)
	}

	products, err := r.getService().Product.FindByCategoryID(context.Background(), uint(categoryID))
	if err != nil {
		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, products)
}

func (r *ProductRest) Show(c *ctx.Context) error {
	idFromUrl := c.GetParam("id")
	id, err := strconv.Atoi(idFromUrl)
	if err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	product, err := r.getService().Product.FindByID(context.Background(), uint(id))
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return response.Fail(c, http.StatusNotFound, err.Error())
		}

		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, product)
}

func (r *ProductRest) Store(c *ctx.Context) error {
	var productIn *request.ProductCreate
	if err := c.GetJSON(&productIn); err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	productToCreate := productIn.ToEntity()
	for _, c := range productIn.Categories {
		cat, err := r.getService().Category.FindByID(context.Background(), c)
		if err != nil {
			continue
		}

		productToCreate.Categories = append(productToCreate.Categories, cat)
	}

	product, err := r.getService().Product.Create(context.Background(), productToCreate)
	if err != nil {
		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, product)
}

func (r *ProductRest) Destroy(c *ctx.Context) error {
	idFromParam := c.GetParam("id")
	id, err := strconv.Atoi(idFromParam)
	if err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	err = r.getService().Product.Delete(context.Background(), uint(id))
	if err != nil {
		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, nil)
}

func (r *ProductRest) GetRatings(c *ctx.Context) error {
	idFromParams := c.GetParam("id")
	id, err := strconv.Atoi(idFromParams)
	if err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	ratings, err := r.getService().Rating.FindByProductID(c.Context(), uint(id))
	if err != nil {
		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, ratings)
}

func (r *ProductRest) SetRatings(c *ctx.Context) error {
	idFromParams := c.GetParam("id")
	id, err := strconv.Atoi(idFromParams)
	if err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	authMiddleware := middleware.NewAuthMiddleware(r.service)
	user, err := authMiddleware.JWT(c)
	if err != nil {
		return response.Fail(c, http.StatusUnauthorized, err.Error())
	}

	var ratingUpdateOrCreateReq *request.RatingUpdateOrCreate
	if c.GetJSON(&ratingUpdateOrCreateReq); err != nil {
		return response.Fail(c, http.StatusBadRequest, err.Error())
	}

	ratingUpdateOrCreateReq.ProductID = uint(id)
	ratingUpdateOrCreateReq.UserID = user.ID

	_, err = r.getService().Rating.Set(c.Context(), ratingUpdateOrCreateReq.ToEntity())
	if err != nil {
		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	product, err := r.getService().Product.FindByID(c.Context(), ratingUpdateOrCreateReq.ProductID)
	if err != nil {
		return response.Fail(c, http.StatusInternalServerError, err.Error())
	}

	return response.Success(c, http.StatusOK, product)
}
