package request

import "github.com/hadihammurabi/dummy-online-shop/app/entity"

type RatingUpdateOrCreate struct {
	ID        uint
	Count     uint `json:"count"`
	UserID    uint
	ProductID uint
}

func (r RatingUpdateOrCreate) ToEntity() *entity.Rating {
	return &entity.Rating{
		ID:        r.ID,
		Count:     r.Count,
		UserID:    r.UserID,
		ProductID: r.ProductID,
	}
}
