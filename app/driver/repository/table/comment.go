package table

import (
	"github.com/hadihammurabi/dummy-online-shop/app/entity"
	"gorm.io/gorm"
)

type Comment struct {
	*gorm.Model
	Content   string
	UserID    uint
	ProductID uint
	CommentID uint
}

func (t *Comment) ToEntity() *entity.Comment {
	c := &entity.Comment{
		ID:      t.ID,
		Content: t.Content,
	}

	return c
}

func CommentFromEntity(t *entity.Comment) *Comment {
	c := &Comment{
		Model: &gorm.Model{
			ID: t.ID,
		},
		Content: t.Content,
	}

	if t.User != nil {
		c.UserID = t.User.ID
	}

	if t.Product != nil {
		c.ProductID = t.Product.ID
	}

	if t.Comment != nil {
		c.CommentID = t.Comment.ID
	}

	return c
}
