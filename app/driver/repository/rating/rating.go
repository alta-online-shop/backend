package rating

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type RatingRepo interface {
	FindByProductID(c context.Context, id uint) ([]entity.Rating, error)
	FindByProductAndUserID(c context.Context, productID, userID uint) (*entity.Rating, error)
	FindByID(c context.Context, id uint) (*entity.Rating, error)
	UpdateByProductAndUserID(c context.Context, productID, userID uint, p *entity.Rating) (*entity.Rating, error)
	CreateByProductAndUserID(c context.Context, productID, userID uint, p *entity.Rating) (*entity.Rating, error)
}
