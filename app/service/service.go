package service

import "github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"

type Service struct {
	Hello   HelloService
	Product ProductService
}

func New() {
	s := &Service{
		Hello:   NewHelloService(),
		Product: NewProductService(),
	}

	ioc.Bind(Service{}, func() interface{} {
		return s
	})
}
