package order

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type OrderRepo interface {
	All(c context.Context) ([]entity.Order, error)
	FindByID(c context.Context, id uint) (*entity.Order, error)
	FindByUserID(c context.Context, id uint) ([]entity.Order, error)
	Create(c context.Context, p *entity.Order) (*entity.Order, error)
	Update(c context.Context, p *entity.Order) (*entity.Order, error)
	Delete(c context.Context, id uint) error
}
