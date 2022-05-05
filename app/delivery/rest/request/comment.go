package request

import "github.com/hadihammurabi/dummy-online-shop/app/entity"

type CommentCreate struct {
	Content   string `json:"content,omitempty"`
	User      *entity.User
	ProductID uint
}

func (r CommentCreate) ToEntity() *entity.Comment {
	return &entity.Comment{
		Content: r.Content,
		User:    r.User,
		Product: &entity.Product{
			ID: r.ProductID,
		},
	}
}
