package product

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type ProductRepo interface {
	All(c context.Context) ([]entity.Product, error)
	FindByID(c context.Context, id uint) (*entity.Product, error)
	Create(c context.Context, p *entity.Product) (*entity.Product, error)
	Update(c context.Context, p *entity.Product) (*entity.Product, error)
	Delete(c context.Context, id uint) error
}
