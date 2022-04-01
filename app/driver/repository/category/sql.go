package category

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type sql struct {
	db *gorm.DB
}

func NewSQL(db *gorm.DB) CategoryRepo {
	return &sql{
		db,
	}
}

func (r *sql) All(c context.Context) ([]entity.Category, error) {
	var categoriesFromTable []Category
	err := r.db.WithContext(c).Find(&categoriesFromTable).Error
	if err != nil {
		return nil, err
	}

	categories := make([]entity.Category, 0)
	for _, p := range categoriesFromTable {
		categories = append(categories, *p.ToEntity())
	}

	return categories, nil
}

func (r *sql) FindByID(c context.Context, id uint) (*entity.Category, error) {
	var categoryFromTable *Category
	err := r.db.WithContext(c).Where("id = ?", id).First(&categoryFromTable).Error
	if err != nil {
		return nil, err
	}

	return categoryFromTable.ToEntity(), nil
}

func (r *sql) Create(c context.Context, p *entity.Category) (*entity.Category, error) {
	categoryToTable := CategoryFromEntity(p)
	err := r.db.WithContext(c).Save(&categoryToTable).Error
	if err != nil {
		return nil, err
	}

	return categoryToTable.ToEntity(), nil
}

func (r *sql) Update(c context.Context, p *entity.Category) (*entity.Category, error) {
	categoryToTable := CategoryFromEntity(p)
	err := r.db.WithContext(c).Updates(&categoryToTable).Error
	if err != nil {
		return nil, err
	}

	return categoryToTable.ToEntity(), nil
}

func (r *sql) Delete(c context.Context, id uint) error {
	err := r.db.WithContext(c).Delete(&Category{Model: &gorm.Model{ID: id}}).Error
	if err != nil {
		return err
	}

	return nil
}
