package product

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type sql struct {
	db *gorm.DB
}

func NewSQL(db *gorm.DB) ProductRepo {
	return &sql{
		db,
	}
}

func (r *sql) All(c context.Context) ([]entity.Product, error) {
	var productsFromTable []Product
	err := r.db.WithContext(c).Find(&productsFromTable).Error
	if err != nil {
		return nil, err
	}

	products := make([]entity.Product, 0)
	for _, p := range productsFromTable {
		products = append(products, p.ToEntity())
	}

	return products, nil
}
