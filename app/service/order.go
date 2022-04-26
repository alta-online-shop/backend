package service

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type OrderService interface {
	All(c context.Context) ([]entity.Order, error)
	FindByID(c context.Context, id uint) (*entity.Order, error)
	Create(c context.Context, p *entity.Order) (*entity.Order, error)
	Update(c context.Context, p *entity.Order) (*entity.Order, error)
	Delete(c context.Context, id uint) error
}

type orderService struct {
	repo *repository.Repository
}

func NewOrderService() OrderService {
	return &orderService{}
}

func (s *orderService) getRepo() *repository.Repository {
	if s.repo == nil {
		s.repo = ioc.Use(repository.Repository{}).(*repository.Repository)
	}

	return s.repo
}

func (s *orderService) All(c context.Context) (orders []entity.Order, err error) {
	orders, err = s.getRepo().Order.All(c)
	return
}

func (s *orderService) FindByID(c context.Context, id uint) (order *entity.Order, err error) {
	order, err = s.getRepo().Order.FindByID(c, id)
	return
}

func (s *orderService) Create(c context.Context, p *entity.Order) (order *entity.Order, err error) {
	order, err = s.getRepo().Order.Create(c, p)
	return
}

func (s *orderService) Update(c context.Context, p *entity.Order) (order *entity.Order, err error) {
	order, err = s.getRepo().Order.Update(c, p)
	return
}

func (s *orderService) Delete(c context.Context, id uint) (err error) {
	err = s.getRepo().Order.Delete(c, id)
	return
}
