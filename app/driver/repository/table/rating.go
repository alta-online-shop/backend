package table

import (
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type Rating struct {
	*gorm.Model
	Count     uint
	UserID    uint
	ProductID uint

	User    *User
	Product *Product
}

func (t *Rating) ToEntity() *entity.Rating {
	c := &entity.Rating{
		ID:    t.ID,
		Count: t.Count,
	}

	if t.User != nil {
		c.User = t.User.ToEntity()
	}

	if t.Product != nil {
		c.Product = t.Product.ToEntity()
	}

	return c
}

func RatingFromEntity(t *entity.Rating) *Rating {
	c := &Rating{
		Model: &gorm.Model{
			ID: t.ID,
		},
		Count: t.Count,
	}

	if t.User != nil {
		c.UserID = t.User.ID
		c.User = UserFromEntity(t.User)
	}

	if t.Product != nil {
		c.ProductID = t.Product.ID
		c.Product = ProductFromEntity(t.Product)
	}

	return c
}
