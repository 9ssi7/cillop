package main

import (
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/env"
	"github.com/mixarchitecture/microp/events/nats"
	"github.com/mixarchitecture/microp/validator"
	"github.com/ssibrahimbas/cillop/config"
	"github.com/ssibrahimbas/cillop/pkg/mongodb"
	event_stream "github.com/ssibrahimbas/cillop/server/event-stream"
	"github.com/ssibrahimbas/cillop/server/http"
	"github.com/ssibrahimbas/cillop/server/rpc"
	"github.com/ssibrahimbas/cillop/service"
)

func main() {
	cnf := config.App{}
	env.Load(&cnf)
	i18n := i18np.New(cnf.I18n.Fallback)
	i18n.Load(cnf.I18n.Dir, cnf.I18n.Locales...)
	valid := validator.New(i18n)
	valid.ConnectCustom()
	valid.RegisterTagName()
	mongo := loadMongo(cnf)
	eventEngine := nats.New(nats.Config{
		Url:     cnf.Nats.Url,
		Streams: cnf.Nats.Streams,
	})
	app := service.NewApp(service.Config{
		App: cnf,
		Db:  mongo,
	})

	http := http.New(http.Config{
		Config: cnf,
		App:    app,
		Valid:  *valid,
		I18n:   i18n,
	})
	rpc := rpc.New(rpc.Config{
		Config: cnf,
		App:    app,
		Valid:  *valid,
		I18n:   *i18n,
	})
	stream := event_stream.New(event_stream.Config{
		Config: cnf,
		App:    app,
		Engine: eventEngine,
	})
	go rpc.Listen()
	go stream.Listen()
	http.Listen()
}

func loadMongo(cnf config.App) *mongodb.DB {
	uri := mongodb.CalcMongoUri(mongodb.UriParams{
		Host:  cnf.ProductMongo.Host,
		Port:  cnf.ProductMongo.Port,
		User:  cnf.ProductMongo.Username,
		Pass:  cnf.ProductMongo.Password,
		Db:    cnf.ProductMongo.Database,
		Query: cnf.ProductMongo.Query,
	})
	d, err := mongodb.New(uri, cnf.ProductMongo.Database)
	if err != nil {
		panic(err)
	}
	return d
}
