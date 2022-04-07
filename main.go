package main

import (
	"log"

	"github.com/hadihammurabi/dummy-online-shop/app/config"
	"github.com/hadihammurabi/dummy-online-shop/app/delivery/web"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/service"
)

func main() {
	config.New()
	if err := repository.New(); err != nil {
		log.Fatalln(err)
	}
	service.New()

	deliveryWeb := web.New()
	log.Fatalln(deliveryWeb.Start())
}
