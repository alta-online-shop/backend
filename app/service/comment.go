package service

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type CommentService interface {
	All(c context.Context) ([]entity.Comment, error)
	FindByProductID(c context.Context, id uint) ([]entity.Comment, error)
	FindByID(c context.Context, id uint) (*entity.Comment, error)
	Create(c context.Context, p *entity.Comment) (*entity.Comment, error)
	Update(c context.Context, p *entity.Comment) (*entity.Comment, error)
	Delete(c context.Context, id uint) error
}

type commentService struct {
	repo *repository.Repository
}

func NewCommentService() CommentService {
	return &commentService{}
}

func (s *commentService) getRepo() *repository.Repository {
	if s.repo == nil {
		s.repo = ioc.Use(repository.Repository{}).(*repository.Repository)
	}

	return s.repo
}

func (s *commentService) All(c context.Context) (comments []entity.Comment, err error) {
	comments, err = s.getRepo().Comment.All(c)
	return
}

func (s *commentService) FindByProductID(c context.Context, id uint) (comments []entity.Comment, err error) {
	comments, err = s.getRepo().Comment.FindByProductID(c, id)
	return
}

func (s *commentService) FindByID(c context.Context, id uint) (comment *entity.Comment, err error) {
	comment, err = s.getRepo().Comment.FindByID(c, id)
	return
}

func (s *commentService) Create(c context.Context, p *entity.Comment) (comment *entity.Comment, err error) {
	comment, err = s.getRepo().Comment.Create(c, p)
	return
}

func (s *commentService) Update(c context.Context, p *entity.Comment) (comment *entity.Comment, err error) {
	comment, err = s.getRepo().Comment.Update(c, p)
	return
}

func (s *commentService) Delete(c context.Context, id uint) (err error) {
	err = s.getRepo().Comment.Delete(c, id)
	return
}
