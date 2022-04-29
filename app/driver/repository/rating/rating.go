package rating

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type RatingRepo interface {
	FindByProductID(c context.Context, id uint) ([]entity.Rating, error)
	FindByID(c context.Context, id uint) (*entity.Rating, error)
	UpdateByProductID(c context.Context, id uint, p *entity.Rating) (*entity.Rating, error)
	CreateByProductID(c context.Context, id uint, p *entity.Rating) (*entity.Rating, error)
}
