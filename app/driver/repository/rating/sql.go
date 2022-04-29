package rating

import (
	"context"
	"errors"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository/table"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type sql struct {
	db *gorm.DB
}

func NewSQL(db *gorm.DB) RatingRepo {
	return &sql{
		db,
	}
}

func (r *sql) FindByProductID(c context.Context, id uint) ([]entity.Rating, error) {
	var ratingsFromTable []table.Rating
	err := r.db.WithContext(c).Where("product_id = ?", id).Preload("User").Preload("Product").Find(&ratingsFromTable).Error
	if err != nil {
		return nil, err
	}

	ratings := make([]entity.Rating, 0)
	for _, p := range ratingsFromTable {
		ratings = append(ratings, *p.ToEntity())
	}

	return ratings, nil
}

func (r *sql) FindByID(c context.Context, id uint) (*entity.Rating, error) {
	var ratingFromTable *table.Rating
	err := r.db.WithContext(c).Where("id = ?", id).First(&ratingFromTable).Error
	if err != nil {
		return nil, err
	}

	return ratingFromTable.ToEntity(), nil
}

func (r *sql) CreateByProductID(c context.Context, id uint, p *entity.Rating) (*entity.Rating, error) {
	ratingToTable := table.RatingFromEntity(p)
	ratingToTable.ProductID = id
	err := r.db.WithContext(c).Create(&ratingToTable).Error
	if err != nil {
		return nil, err
	}
	return ratingToTable.ToEntity(), nil
}

func (r *sql) UpdateByProductID(c context.Context, id uint, p *entity.Rating) (*entity.Rating, error) {
	ratingsFromTable, err := r.FindByProductID(c, id)
	if err != nil {
		return nil, err
	}

	if len(ratingsFromTable) <= 0 {
		return nil, errors.New("not found")
	}

	ratingToTable := table.RatingFromEntity(&ratingsFromTable[0])
	ratingToTable.Count = p.Count
	err = r.db.WithContext(c).Updates(&ratingToTable).Error
	if err != nil {
		return nil, err
	}

	return ratingToTable.ToEntity(), nil
}
