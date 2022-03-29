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
}

func (t *Product) ToEntity() entity.Product {
	return entity.Product{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Price:       t.Price,
	}
}
