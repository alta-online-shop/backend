package table

import (
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model
	Name        string
	Description string
	Price       uint

	Categories []*Category `gorm:"many2many:product_categories;"`
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

	product.Categories = make([]*Category, 0)
	for _, cat := range t.Categories {
		product.Categories = append(product.Categories, CategoryFromEntity(cat))
	}

	return product
}
