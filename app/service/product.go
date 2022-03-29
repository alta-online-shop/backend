package service

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type ProductService interface {
	All(c context.Context) ([]entity.Product, error)
}

type productService struct {
	repo *repository.Repository
}

func NewProductService() ProductService {
	return &productService{}
}

func (s *productService) getRepo() *repository.Repository {
	if s.repo == nil {
		s.repo = ioc.Use(repository.Repository{}).(*repository.Repository)
	}

	return s.repo
}

func (s *productService) All(c context.Context) (products []entity.Product, err error) {
	products, err = s.getRepo().Product.All(c)
	return
}
