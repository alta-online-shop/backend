package user

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type UserRepo interface {
	All(c context.Context) ([]entity.User, error)
	FindByID(c context.Context, id uint) (*entity.User, error)
	FindByEmail(c context.Context, email string) (*entity.User, error)
	Create(c context.Context, p *entity.User) (*entity.User, error)
	Update(c context.Context, p *entity.User) (*entity.User, error)
	Delete(c context.Context, id uint) error
}
