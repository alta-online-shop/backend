package service

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
)

type RatingService interface {
	FindByProductID(c context.Context, id uint) (uint, error)
	FindByProductAndUserID(c context.Context, productID, userID uint) (*entity.Rating, error)
	FindByID(c context.Context, id uint) (*entity.Rating, error)
	Set(c context.Context, p *entity.Rating) (*entity.Rating, error)
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

	length := uint(len(ratings))
	if length > 0 {
		count /= uint(len(ratings))
	}

	return
}

func (s *ratingService) FindByProductAndUserID(c context.Context, productID, userID uint) (rating *entity.Rating, err error) {
	rating, err = s.getRepo().Rating.FindByProductAndUserID(c, productID, userID)
	return
}

func (s *ratingService) FindByID(c context.Context, id uint) (rating *entity.Rating, err error) {
	rating, err = s.getRepo().Rating.FindByID(c, id)
	return
}

func (s *ratingService) Set(c context.Context, p *entity.Rating) (rating *entity.Rating, err error) {
	_, err = s.FindByProductAndUserID(c, p.ProductID, p.UserID)
	if err != nil {
		if rating, err = s.getRepo().Rating.CreateByProductAndUserID(c, p.ProductID, p.UserID, p); err != nil {
			return
		}
	} else {
		if rating, err = s.getRepo().Rating.UpdateByProductAndUserID(c, p.ProductID, p.UserID, p); err != nil {
			return
		}
	}

	return
}
