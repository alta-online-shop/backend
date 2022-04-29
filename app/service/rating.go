package service

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type RatingService interface {
	FindByProductID(c context.Context, id uint) (uint, error)
	FindByID(c context.Context, id uint) (*entity.Rating, error)
	UpdateOrCreate(c context.Context, id uint, p *entity.Rating) (*entity.Rating, error)
}

type ratingService struct {
	repo *repository.Repository
}

func NewRatingService() RatingService {
	return &ratingService{}
}

func (s *ratingService) getRepo() *repository.Repository {
	if s.repo == nil {
		s.repo = ioc.Use(repository.Repository{}).(*repository.Repository)
	}

	return s.repo
}

func (s *ratingService) FindByProductID(c context.Context, id uint) (count uint, err error) {
	ratings, err := s.getRepo().Rating.FindByProductID(c, id)
	if err != nil {
		return
	}

	for _, rating := range ratings {
		count += rating.Count
	}

	return
}

func (s *ratingService) FindByID(c context.Context, id uint) (rating *entity.Rating, err error) {
	rating, err = s.getRepo().Rating.FindByID(c, id)
	return
}

func (s *ratingService) UpdateOrCreate(c context.Context, id uint, p *entity.Rating) (rating *entity.Rating, err error) {
	ratings, err := s.getRepo().Rating.FindByProductID(c, id)
	if err != nil {
		return
	}

	if len(ratings) <= 0 {
		rating, err = s.getRepo().Rating.CreateByProductID(c, id, p)
		if err != nil {
			return nil, err
		}
	} else {
		rating, err = s.getRepo().Rating.UpdateByProductID(c, id, p)
		if err != nil {
			return
		}
	}

	product, _ := s.getRepo().Product.FindByID(c, id)
	rating.Product = product

	return
}
