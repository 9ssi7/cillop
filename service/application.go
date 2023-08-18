package service

import (
	"github.com/ssibrahimbas/cillop/app"
	"github.com/ssibrahimbas/cillop/app/command"
	"github.com/ssibrahimbas/cillop/app/query"
	"github.com/ssibrahimbas/cillop/config"
	"github.com/ssibrahimbas/cillop/pkg/domains/product"
	"github.com/ssibrahimbas/cillop/pkg/mongodb"
)

type Config struct {
	App config.App
	Db  *mongodb.DB
}

func NewApp(config Config) app.Application {
	productRepo := product.NewRepository(config.Db.GetCollection(config.App.ProductMongo.Collection))

	return app.Application{
		Commands: app.Commands{
			ProductCreate: command.NewProductCreateHandler(productRepo),
		},
		Queries: app.Queries{
			ProductGet: query.NewProductGetHandler(productRepo),
		},
	}
}
