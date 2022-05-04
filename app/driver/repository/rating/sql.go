package rating

import (
	"context"

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

func (r *sql) FindByProductAndUserID(c context.Context, productID, userID uint) (rating *entity.Rating, err error) {
	var ratingFromTable *table.Rating
	err = r.db.WithContext(c).
		Preload("User", "id = ?", userID).
		Preload("Product").
		Where("product_id = ?", productID).
		First(&ratingFromTable).Error
	if err != nil {
		return nil, err
	}

	rating = ratingFromTable.ToEntity()
	return rating, nil
}

func (r *sql) FindByID(c context.Context, id uint) (*entity.Rating, error) {
	var ratingFromTable *table.Rating
	err := r.db.WithContext(c).Where("id = ?", id).First(&ratingFromTable).Error
	if err != nil {
		return nil, err
	}

	return ratingFromTable.ToEntity(), nil
}

func (r *sql) CreateByProductAndUserID(c context.Context, productID, userID uint, p *entity.Rating) (*entity.Rating, error) {
	ratingToTable := table.RatingFromEntity(p)
	ratingToTable.ProductID = productID
	ratingToTable.UserID = userID
	err := r.db.WithContext(c).Create(&ratingToTable).Error
	if err != nil {
		return nil, err
	}
	return ratingToTable.ToEntity(), nil
}

func (r *sql) UpdateByProductAndUserID(c context.Context, productID, userID uint, p *entity.Rating) (*entity.Rating, error) {
	ratingsFromTable, err := r.FindByProductAndUserID(c, productID, userID)
	if err != nil {
		return nil, err
	}

	ratingToTable := table.RatingFromEntity(ratingsFromTable)
	ratingToTable.Count = p.Count
	err = r.db.Debug().WithContext(c).Updates(&ratingToTable).Error
	if err != nil {
		return nil, err
	}

	return ratingToTable.ToEntity(), nil
}
