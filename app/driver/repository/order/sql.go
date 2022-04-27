package order

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository/product"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository/table"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository/user"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type sql struct {
	db          *gorm.DB
	userRepo    user.UserRepo
	productRepo product.ProductRepo
}

func NewSQL(db *gorm.DB, userRepo user.UserRepo, productRepo product.ProductRepo) OrderRepo {
	return &sql{
		db,
		userRepo,
		productRepo,
	}
}

func (r *sql) All(c context.Context) ([]entity.Order, error) {
	var orderFromTable []table.Order
	err := r.db.WithContext(c).Find(&orderFromTable).Error
	if err != nil {
		return nil, err
	}

	orders := make([]entity.Order, 0)
	for _, p := range orderFromTable {
		user, _ := r.userRepo.FindByID(c, p.UserID)
		product, _ := r.productRepo.FindByID(c, p.ProductID)
		orders = append(orders, *p.ToEntity(user, product))
	}

	return orders, nil
}

func (r *sql) FindByID(c context.Context, id uint) (*entity.Order, error) {
	var orderFromTable *table.Order
	err := r.db.WithContext(c).Where("id = ?", id).First(&orderFromTable).Error
	if err != nil {
		return nil, err
	}

	user, _ := r.userRepo.FindByID(c, orderFromTable.UserID)
	product, _ := r.productRepo.FindByID(c, orderFromTable.ProductID)
	return orderFromTable.ToEntity(user, product), nil
}

func (r *sql) FindByUserID(c context.Context, id uint) ([]entity.Order, error) {
	var orderFromTable []table.Order
	err := r.db.WithContext(c).Where("user_id = ?", id).Find(&orderFromTable).Error
	if err != nil {
		return nil, err
	}

	orders := make([]entity.Order, 0)
	for _, p := range orderFromTable {
		user, _ := r.userRepo.FindByID(c, p.UserID)
		product, _ := r.productRepo.FindByID(c, p.ProductID)
		orders = append(orders, *p.ToEntity(user, product))
	}

	return orders, nil
}

func (r *sql) Create(c context.Context, p *entity.Order) (*entity.Order, error) {
	orderToTable := table.OrderFromEntity(p)
	err := r.db.WithContext(c).Create(&orderToTable).Error
	if err != nil {
		return nil, err
	}

	user, _ := r.userRepo.FindByID(c, orderToTable.UserID)
	product, _ := r.productRepo.FindByID(c, orderToTable.ProductID)
	return orderToTable.ToEntity(user, product), nil
}

func (r *sql) Update(c context.Context, p *entity.Order) (*entity.Order, error) {
	orderToTable := table.OrderFromEntity(p)
	err := r.db.WithContext(c).Updates(&orderToTable).Error
	if err != nil {
		return nil, err
	}
	user, _ := r.userRepo.FindByID(c, orderToTable.UserID)
	product, _ := r.productRepo.FindByID(c, orderToTable.ProductID)
	return orderToTable.ToEntity(user, product), nil
}

func (r *sql) Delete(c context.Context, id uint) error {
	err := r.db.WithContext(c).Delete(&table.Order{Model: &gorm.Model{ID: id}}).Error
	if err != nil {
		return err
	}

	return nil
}
