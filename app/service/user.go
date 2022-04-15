package service

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type UserService interface {
	All(c context.Context) ([]entity.User, error)
	FindByID(c context.Context, id uint) (*entity.User, error)
	FindByEmail(c context.Context, email string) (*entity.User, error)
	Create(c context.Context, p *entity.User) (*entity.User, error)
	Update(c context.Context, p *entity.User) (*entity.User, error)
	Delete(c context.Context, id uint) error
}

type userService struct {
	repo *repository.Repository
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) getRepo() *repository.Repository {
	if s.repo == nil {
		s.repo = ioc.Use(repository.Repository{}).(*repository.Repository)
	}

	return s.repo
}

func (s *userService) All(c context.Context) (categories []entity.User, err error) {
	categories, err = s.getRepo().User.All(c)
	return
}

func (s *userService) FindByID(c context.Context, id uint) (user *entity.User, err error) {
	user, err = s.getRepo().User.FindByID(c, id)
	return
}

func (s *userService) FindByEmail(c context.Context, email string) (user *entity.User, err error) {
	user, err = s.getRepo().User.FindByEmail(c, email)
	return
}

func (s *userService) Create(c context.Context, p *entity.User) (user *entity.User, err error) {
	user, err = s.getRepo().User.Create(c, p)
	return
}

func (s *userService) Update(c context.Context, p *entity.User) (user *entity.User, err error) {
	user, err = s.getRepo().User.Update(c, p)
	return
}

func (s *userService) Delete(c context.Context, id uint) (err error) {
	err = s.getRepo().User.Delete(c, id)
	return
}
