package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/server/http"
	"github.com/mixarchitecture/microp/server/http/parser"
	"github.com/mixarchitecture/microp/validator"
	"github.com/ssibrahimbas/cillop/app"
	"github.com/ssibrahimbas/cillop/config"
	"github.com/ssibrahimbas/cillop/pkg/server"
)

type srv struct {
	config config.App
	app    app.Application
	valid  validator.Validator
	i18n   *i18np.I18n
}

type Config struct {
	Config config.App
	App    app.Application
	Valid  validator.Validator
	I18n   *i18np.I18n
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
	http.RunServer(http.Config{
		Host: s.config.Http.Host,
		Port: s.config.Http.Port,
		I18n: s.i18n,
		CreateHandler: func(router fiber.Router) fiber.Router {
			router.Post("/", s.ProductCreate)
			router.Get("/:id", s.ProductGet)
			return router
		},
	})
	return nil
}

func (s *srv) parseBody(ctx *fiber.Ctx, dto interface{}) {
	parser.ParseBody(ctx, s.valid, *s.i18n, dto)
}

func (s *srv) parseParams(ctx *fiber.Ctx, dto interface{}) {
	parser.ParseParams(ctx, s.valid, *s.i18n, dto)
}

func (s *srv) parseQuery(ctx *fiber.Ctx, dto interface{}) {
	parser.ParseQuery(ctx, s.valid, *s.i18n, dto)
}
