package table

import (
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type Order struct {
	*gorm.Model
	UserID    uint
	ProductID uint
	Quantity  uint
}

func (t *Order) ToEntity(user *entity.User, product *entity.Product) *entity.Order {
	order := &entity.Order{
		ID:       t.ID,
		User:     user,
		Product:  product,
		Quantity: t.Quantity,
	}

	return order
}

func OrderFromEntity(t *entity.Order) *Order {
	order := &Order{
		Model: &gorm.Model{
			ID: t.ID,
		},
		UserID:    t.User.ID,
		ProductID: t.Product.ID,
		Quantity:  t.Quantity,
	}

	return order
}
