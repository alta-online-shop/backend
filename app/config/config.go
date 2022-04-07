package config

import (
	"os"

	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
)

type Config struct {
	DBType      string
	DBDSN       string
	ResourceDir string
}

func New() {
	dbtype := os.Getenv("APP_DB_TYPE")
	if dbtype == "" {
		dbtype = "pg"
	}

	dbdsn := os.Getenv("APP_DB_DSN")

	resourceDir := os.Getenv("APP_RESOURCE_DIR")
	if resourceDir == "" {
		resourceDir = "resources"
	}

	c := &Config{
		DBType:      dbtype,
		DBDSN:       dbdsn,
		ResourceDir: resourceDir,
	}

	ioc.Bind(Config{}, func() interface{} {
		return c
	})
}
