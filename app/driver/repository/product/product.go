package product

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model
	Name        string
	Description string
	Price       uint
}

type ProductRepo interface {
	All(c context.Context) ([]entity.Product, error)
	FindByID(c context.Context, id uint) (*entity.Product, error)
	Create(c context.Context, p *entity.Product) (*entity.Product, error)
	Update(c context.Context, p *entity.Product) (*entity.Product, error)
	Delete(c context.Context, id uint) error
}

func (t *Product) ToEntity() *entity.Product {
	return &entity.Product{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Price:       t.Price,
	}
}

func ProductFromEntity(t *entity.Product) *Product {
	return &Product{
		Model: &gorm.Model{
			ID: t.ID,
		},
		Name:        t.Name,
		Description: t.Description,
		Price:       t.Price,
	}
}
