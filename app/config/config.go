package config

import (
	"os"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
)

type Config struct {
	DBType string
	DBDSN  string
}

func New() {
	dbtype := os.Getenv("APP_DB_TYPE")
	dbdsn := os.Getenv("APP_DB_DSN")

	c := &Config{
		DBType: dbtype,
		DBDSN:  dbdsn,
	}

	ioc.Bind(Config{}, func() interface{} {
		return c
	})
}
