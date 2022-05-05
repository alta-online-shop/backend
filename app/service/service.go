package service

import "github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"

type Service struct {
	Hello    HelloService
	Product  ProductService
	Category CategoryService
	User     UserService
	Auth     AuthService
	Order    OrderService
	Rating   RatingService
	Comment  CommentService
}

func New() {
	s := &Service{
		Hello:    NewHelloService(),
		Product:  NewProductService(),
		Category: NewCategoryService(),
		User:     NewUserService(),
		Auth:     NewAuthService(),
		Order:    NewOrderService(),
		Rating:   NewRatingService(),
		Comment:  NewCommentService(),
	}

	ioc.Bind(Service{}, func() interface{} {
		return s
	})
}
