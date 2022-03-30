package repository

import (
	"errors"

	"github.com/hadihammurabi/dummy-online-shop/app/config"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/ioc"
	"github.com/hadihammurabi/dummy-online-shop/app/driver/repository/product"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	Product product.ProductRepo
}

func builldRepo(db *gorm.DB) *Repository {
	return &Repository{
		Product: product.NewSQL(db),
	}
}

func New() error {
	cfg := ioc.Use(config.Config{}).(*config.Config)

	db, err := setup(cfg)
	if err != nil {
		return err
	}
	repo := builldRepo(db)

	ioc.Bind(Repository{}, func() interface{} {
		return repo
	})
	return nil
}

func setup(cfg *config.Config) (*gorm.DB, error) {
	if cfg.DBType == "pg" {
		db, err := gorm.Open(postgres.Open(cfg.DBDSN))
		return db, err
	}

	if cfg.DBType == "mysql" {
		db, err := gorm.Open(mysql.Open(cfg.DBDSN))
		return db, err
	}

	return nil, errors.New("unknown database type")
}
