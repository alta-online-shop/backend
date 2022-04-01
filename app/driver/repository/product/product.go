package product

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository/category"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model
	Name        string
	Description string
	Price       uint

	Categories []*category.Category `gorm:"many2many:product_categories;"`
}

type ProductRepo interface {
	All(c context.Context) ([]entity.Product, error)
	FindByID(c context.Context, id uint) (*entity.Product, error)
	Create(c context.Context, p *entity.Product) (*entity.Product, error)
	Update(c context.Context, p *entity.Product) (*entity.Product, error)
	Delete(c context.Context, id uint) error
}

func (t *Product) ToEntity() *entity.Product {
	product := &entity.Product{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Price:       t.Price,
	}

	product.Categories = make([]*entity.Category, 0)
	for _, cat := range t.Categories {
		product.Categories = append(product.Categories, cat.ToEntity())
	}

	return product
}

func ProductFromEntity(t *entity.Product) *Product {
	product := &Product{
		Model: &gorm.Model{
			ID: t.ID,
		},
		Name:        t.Name,
		Description: t.Description,
		Price:       t.Price,
	}

	product.Categories = make([]*category.Category, 0)
	for _, cat := range t.Categories {
		product.Categories = append(product.Categories, category.CategoryFromEntity(cat))
	}

	return product
}
