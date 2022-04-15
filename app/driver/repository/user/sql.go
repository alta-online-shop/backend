package user

import (
	"context"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository/table"
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type sql struct {
	db *gorm.DB
}

func NewSQL(db *gorm.DB) UserRepo {
	return &sql{
		db,
	}
}

func (r *sql) All(c context.Context) ([]entity.User, error) {
	var usersFromTable []table.User
	err := r.db.WithContext(c).Find(&usersFromTable).Error
	if err != nil {
		return nil, err
	}

	users := make([]entity.User, 0)
	for _, p := range usersFromTable {
		users = append(users, *p.ToEntity())
	}

	return users, nil
}

func (r *sql) FindByID(c context.Context, id uint) (*entity.User, error) {
	var userFromTable *table.User
	err := r.db.WithContext(c).Where("id = ?", id).First(&userFromTable).Error
	if err != nil {
		return nil, err
	}

	return userFromTable.ToEntity(), nil
}

func (r *sql) FindByEmail(c context.Context, email string) (*entity.User, error) {
	var userFromTable *table.User
	err := r.db.WithContext(c).Where("email = ?", email).First(&userFromTable).Error
	if err != nil {
		return nil, err
	}

	return userFromTable.ToEntity(), nil
}

func (r *sql) Create(c context.Context, p *entity.User) (*entity.User, error) {
	userToTable := table.UserFromEntity(p)
	err := r.db.WithContext(c).Create(&userToTable).Error
	if err != nil {
		return nil, err
	}

	return userToTable.ToEntity(), nil
}

func (r *sql) Update(c context.Context, p *entity.User) (*entity.User, error) {
	userToTable := table.UserFromEntity(p)
	err := r.db.WithContext(c).Updates(&userToTable).Error
	if err != nil {
		return nil, err
	}

	return userToTable.ToEntity(), nil
}

func (r *sql) Delete(c context.Context, id uint) error {
	err := r.db.WithContext(c).Delete(&table.User{Model: &gorm.Model{ID: id}}).Error
	if err != nil {
		return err
	}

	return nil
}
