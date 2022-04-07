package category

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type CategoryRepo interface {
	All(c context.Context) ([]entity.Category, error)
	FindByID(c context.Context, id uint) (*entity.Category, error)
	Create(c context.Context, p *entity.Category) (*entity.Category, error)
	Update(c context.Context, p *entity.Category) (*entity.Category, error)
	Delete(c context.Context, id uint) error
}
