package rpc

import (
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/server/rpc"
	"github.com/mixarchitecture/microp/validator"
	"github.com/ssibrahimbas/cillop/app"
	"github.com/ssibrahimbas/cillop/config"
	"github.com/ssibrahimbas/cillop/pkg/server"
	"github.com/ssibrahimbas/cillop/server/rpc/routes"
	"google.golang.org/grpc"
)

type srv struct {
	routes.ProductServiceServer
	config config.App
	app    app.Application
	valid  validator.Validator
	i18n   i18np.I18n
}

type Config struct {
	Config config.App
	App    app.Application
	Valid  validator.Validator
	I18n   i18np.I18n
}

func New(cfg Config) server.Server {
	return &srv{
		config: cfg.Config,
		app:    cfg.App,
		valid:  cfg.Valid,
		i18n:   cfg.I18n,
	}
}

func (s *srv) Listen() error {
	rpc.RunServer(s.config.Rpc.Port, func(server *grpc.Server) {
		routes.RegisterProductServiceServer(server, s)
	})
	return nil
}
