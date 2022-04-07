package table

import (
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type Category struct {
	*gorm.Model
	Name        string
	Description string
	Products    []*Product `gorm:"many2many:product_categories;"`
}

func (t *Category) ToEntity() *entity.Category {
	return &entity.Category{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
	}
}

func CategoryFromEntity(t *entity.Category) *Category {
	return &Category{
		Model: &gorm.Model{
			ID: t.ID,
		},
		Name:        t.Name,
		Description: t.Description,
	}
}
