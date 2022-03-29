package main

import (
	"log"

	"github.com/hadihammurabi/dummy-online-shop/app/api/rest"
	"github.com/hadihammurabi/dummy-online-shop/app/config"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository"
	"github.com/hadihammurabi/dummy-online-shop/app/service"
)

func main() {
	config.New()
	if err := repository.New(); err != nil {
		log.Fatalln(err)
	}
	service.New()

	apiRest := rest.New()
	log.Fatalln(apiRest.Start())
}
