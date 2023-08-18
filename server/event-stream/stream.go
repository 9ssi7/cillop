package event_stream

import (
	"github.com/mixarchitecture/microp/events"
	"github.com/ssibrahimbas/cillop/app"
	"github.com/ssibrahimbas/cillop/config"
	"github.com/ssibrahimbas/cillop/pkg/server"
)

type srv struct {
	config config.App
	app    app.Application
	engine events.Engine
}

type Config struct {
	Config config.App
	App    app.Application
	Engine events.Engine
}

func New(cfg Config) server.Server {
	return &srv{
		config: cfg.Config,
		app:    cfg.App,
		engine: cfg.Engine,
	}
}

func (s *srv) Listen() error {
	return s.engine.Subscribe(s.config.EventStream.TopicProductCreate, s.CreateProduct)
}
