package product

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository/table"
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
	var productsFromTable []table.Product
	err := r.db.Debug().WithContext(c).Preload("Categories").Find(&productsFromTable).Error
	if err != nil {
		return nil, err
	}

	products := make([]entity.Product, 0)
	for _, p := range productsFromTable {
		products = append(products, *p.ToEntity())
	}

	return products, nil
}

func (r *sql) FindByID(c context.Context, id uint) (*entity.Product, error) {
	var productFromTable *table.Product
	err := r.db.WithContext(c).Where("id = ?", id).Preload("Categories").First(&productFromTable).Error
	if err != nil {
		return nil, err
	}

	return productFromTable.ToEntity(), nil
}

func (r *sql) Create(c context.Context, p *entity.Product) (*entity.Product, error) {
	productToTable := table.ProductFromEntity(p)
	err := r.db.Debug().WithContext(c).Create(&productToTable).Error
	if err != nil {
		return nil, err
	}

	return productToTable.ToEntity(), nil
}

func (r *sql) Update(c context.Context, p *entity.Product) (*entity.Product, error) {
	productToTable := table.ProductFromEntity(p)
	err := r.db.WithContext(c).Updates(&productToTable).Error
	if err != nil {
		return nil, err
	}

	return productToTable.ToEntity(), nil
}

func (r *sql) Delete(c context.Context, id uint) error {
	err := r.db.WithContext(c).Delete(&table.Product{Model: &gorm.Model{ID: id}}).Error
	if err != nil {
		return err
	}

	return nil
}
