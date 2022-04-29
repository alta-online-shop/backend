package request

import "github.com/hadihammurabi/dummy-online-shop/app/entity"

type RatingUpdateOrCreate struct {
	Count uint `json:"count"`
	User  *entity.User
}

func (r RatingUpdateOrCreate) ToEntity() *entity.Rating {
	return &entity.Rating{
		Count: r.Count,
		User:  r.User,
	}
}
