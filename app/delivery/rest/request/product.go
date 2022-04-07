package request

import "github.com/hadihammurabi/dummy-online-shop/app/entity"

type ProductCreate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Price       uint   `json:"price"`
	Categories  []uint `json:"categories,omitempty"`
}

func (r ProductCreate) ToEntity() *entity.Product {
	return &entity.Product{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
	}
}
