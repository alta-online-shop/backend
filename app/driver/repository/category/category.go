package category

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type Category struct {
	*gorm.Model
	Name        string
	Description string
}

type CategoryRepo interface {
	All(c context.Context) ([]entity.Category, error)
	FindByID(c context.Context, id uint) (*entity.Category, error)
	Create(c context.Context, p *entity.Category) (*entity.Category, error)
	Update(c context.Context, p *entity.Category) (*entity.Category, error)
	Delete(c context.Context, id uint) error
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
