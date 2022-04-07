package request

import "github.com/hadihammurabi/dummy-online-shop/app/entity"

type CategoryCreate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r CategoryCreate) ToEntity() *entity.Category {
	return &entity.Category{
		Name:        r.Name,
		Description: r.Description,
	}
}
