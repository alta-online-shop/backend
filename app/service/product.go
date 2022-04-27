package service

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type ProductService interface {
	All(c context.Context) ([]entity.Product, error)
	FindByCategoryID(c context.Context, id uint) ([]entity.Product, error)
	FindByID(c context.Context, id uint) (*entity.Product, error)
	Create(c context.Context, p *entity.Product) (*entity.Product, error)
	Update(c context.Context, p *entity.Product) (*entity.Product, error)
	Delete(c context.Context, id uint) error
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

func (s *productService) FindByCategoryID(c context.Context, id uint) (products []entity.Product, err error) {
	products, err = s.getRepo().Product.FindByCategoryID(c, id)
	return
}

func (s *productService) FindByID(c context.Context, id uint) (product *entity.Product, err error) {
	product, err = s.getRepo().Product.FindByID(c, id)
	return
}

func (s *productService) Create(c context.Context, p *entity.Product) (product *entity.Product, err error) {
	product, err = s.getRepo().Product.Create(c, p)
	return
}

func (s *productService) Update(c context.Context, p *entity.Product) (product *entity.Product, err error) {
	product, err = s.getRepo().Product.Update(c, p)
	return
}

func (s *productService) Delete(c context.Context, id uint) (err error) {
	err = s.getRepo().Product.Delete(c, id)
	return
}
