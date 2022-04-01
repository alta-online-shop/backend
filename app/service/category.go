package service

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type CategoryService interface {
	All(c context.Context) ([]entity.Category, error)
	FindByID(c context.Context, id uint) (*entity.Category, error)
	Create(c context.Context, p *entity.Category) (*entity.Category, error)
	Update(c context.Context, p *entity.Category) (*entity.Category, error)
	Delete(c context.Context, id uint) error
}

type categoryService struct {
	repo *repository.Repository
}

func NewCategoryService() CategoryService {
	return &categoryService{}
}

func (s *categoryService) getRepo() *repository.Repository {
	if s.repo == nil {
		s.repo = ioc.Use(repository.Repository{}).(*repository.Repository)
	}

	return s.repo
}

func (s *categoryService) All(c context.Context) (categories []entity.Category, err error) {
	categories, err = s.getRepo().Category.All(c)
	return
}

func (s *categoryService) FindByID(c context.Context, id uint) (category *entity.Category, err error) {
	category, err = s.getRepo().Category.FindByID(c, id)
	return
}

func (s *categoryService) Create(c context.Context, p *entity.Category) (category *entity.Category, err error) {
	category, err = s.getRepo().Category.Create(c, p)
	return
}

func (s *categoryService) Update(c context.Context, p *entity.Category) (category *entity.Category, err error) {
	category, err = s.getRepo().Category.Update(c, p)
	return
}

func (s *categoryService) Delete(c context.Context, id uint) (err error) {
	err = s.getRepo().Category.Delete(c, id)
	return
}
