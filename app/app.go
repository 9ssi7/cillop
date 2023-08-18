package app

import (
	"github.com/ssibrahimbas/cillop/app/command"
	"github.com/ssibrahimbas/cillop/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ProductCreate command.ProductCreateHandler
}

type Queries struct {
	ProductGet query.ProductGetHandler
}
