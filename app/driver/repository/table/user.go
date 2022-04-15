package table

import (
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Fullname string
	Email    string
	Password string
}

func (t *User) ToEntity() *entity.User {
	user := &entity.User{
		ID:       t.ID,
		Fullname: t.Fullname,
		Email:    t.Email,
		Password: t.Password,
	}

	return user
}

func UserFromEntity(t *entity.User) *User {
	user := &User{
		Model: &gorm.Model{
			ID: t.ID,
		},
		Fullname: t.Fullname,
		Email:    t.Email,
		Password: t.Password,
	}

	return user
}
