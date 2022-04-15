package service

import "github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"

type Service struct {
	Hello    HelloService
	Product  ProductService
	Category CategoryService
	User     UserService
}

func New() {
	s := &Service{
		Hello:    NewHelloService(),
		Product:  NewProductService(),
		Category: NewCategoryService(),
		User:     NewUserService(),
	}

	ioc.Bind(Service{}, func() interface{} {
		return s
	})
}
