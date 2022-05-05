package comment

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type CommentRepo interface {
	All(c context.Context) ([]entity.Comment, error)
	FindByProductID(c context.Context, id uint) ([]entity.Comment, error)
	FindByID(c context.Context, id uint) (*entity.Comment, error)
	Create(c context.Context, p *entity.Comment) (*entity.Comment, error)
	Update(c context.Context, p *entity.Comment) (*entity.Comment, error)
	Delete(c context.Context, id uint) error
}
