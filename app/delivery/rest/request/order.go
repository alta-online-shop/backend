package request

import "github.com/hadihammurabi/dummy-online-shop/app/entity"

type OrderCreate struct {
	ProductID uint `json:"product_id,omitempty"`
	Quantity  uint `json:"quantity,omitempty"`
}

func (r OrderCreate) ToEntity() *entity.Order {
	return &entity.Order{
		Product: &entity.Product{
			ID: r.ProductID,
		},
		Quantity: r.Quantity,
	}
}
