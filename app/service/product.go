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
	sv   *Service
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

func (s *productService) getService() *Service {
	if s.sv == nil {
		s.sv = ioc.Use(Service{}).(*Service)
	}

	return s.sv
}

func (s *productService) All(c context.Context) (products []entity.Product, err error) {
	currentProducts, err := s.getRepo().Product.All(c)
	if err != nil {
		return
	}

	for _, product := range currentProducts {
		rating, err := s.getService().Rating.FindByProductID(c, product.ID)
		if err != nil {
			continue
		}

		product.Ratings = rating
		products = append(products, product)
	}

	return
}

func (s *productService) FindByCategoryID(c context.Context, id uint) (products []entity.Product, err error) {
	products = make([]entity.Product, 0)
	currentProducts, err := s.getRepo().Product.FindByCategoryID(c, id)
	if err != nil {
		return
	}

	for _, product := range currentProducts {
		rating, err := s.getService().Rating.FindByProductID(c, product.ID)
		if err != nil {
			continue
		}

		product.Ratings = rating
		products = append(products, product)
	}

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
