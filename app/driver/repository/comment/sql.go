package comment

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository/table"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type sql struct {
	db *gorm.DB
}

func NewSQL(db *gorm.DB) CommentRepo {
	return &sql{
		db,
	}
}

func (r *sql) All(c context.Context) ([]entity.Comment, error) {
	var categoriesFromTable []table.Comment
	err := r.db.WithContext(c).Find(&categoriesFromTable).Error
	if err != nil {
		return nil, err
	}

	categories := make([]entity.Comment, 0)
	for _, p := range categoriesFromTable {
		categories = append(categories, *p.ToEntity())
	}

	return categories, nil
}

func (r *sql) FindByProductID(c context.Context, id uint) ([]entity.Comment, error) {
	var categoriesFromTable []table.Comment
	err := r.db.WithContext(c).Where("product_id = ?", id).Find(&categoriesFromTable).Error
	if err != nil {
		return nil, err
	}

	categories := make([]entity.Comment, 0)
	for _, p := range categoriesFromTable {
		categories = append(categories, *p.ToEntity())
	}

	return categories, nil
}

func (r *sql) FindByID(c context.Context, id uint) (*entity.Comment, error) {
	var categoryFromTable *table.Comment
	err := r.db.WithContext(c).Where("id = ?", id).First(&categoryFromTable).Error
	if err != nil {
		return nil, err
	}

	return categoryFromTable.ToEntity(), nil
}

func (r *sql) Create(c context.Context, p *entity.Comment) (*entity.Comment, error) {
	commentToTable := table.CommentFromEntity(p)
	err := r.db.WithContext(c).Create(&commentToTable).Error
	if err != nil {
		return nil, err
	}

	return commentToTable.ToEntity(), nil
}

func (r *sql) Update(c context.Context, p *entity.Comment) (*entity.Comment, error) {
	commentToTable := table.CommentFromEntity(p)
	err := r.db.WithContext(c).Updates(&commentToTable).Error
	if err != nil {
		return nil, err
	}

	return commentToTable.ToEntity(), nil
}

func (r *sql) Delete(c context.Context, id uint) error {
	err := r.db.WithContext(c).Delete(&table.Comment{Model: &gorm.Model{ID: id}}).Error
	if err != nil {
		return err
	}

	return nil
}
