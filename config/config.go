package config

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type HttpServer struct {
	Host string `env:"HTTP_SERVER_HOST" envDefault:"localhost"`
	Port int    `env:"HTTP_SERVER_PORT" envDefault:"3000"`
}

type RpcServer struct {
	Port int `env:"RPC_SERVER_PORT" envDefault:"3001"`
}

type EventStream struct {
	TopicProductCreate string `env:"EVENT_STREAM_TOPIC_PRODUCT_CREATE" envDefault:"product.create"`
}

type ProductMongo struct {
	Host       string `env:"MONGO_PRODUCT_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_PRODUCT_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_PRODUCT_USERNAME" envDefault:""`
	Password   string `env:"MONGO_PRODUCT_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_PRODUCT_DATABASE" envDefault:"product"`
	Collection string `env:"MONGO_PRODUCT_COLLECTION" envDefault:"products"`
	Query      string `env:"MONGO_PRODUCT_QUERY" envDefault:""`
}

type NatsConfig struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type App struct {
	I18n         I18n
	Http         HttpServer
	ProductMongo ProductMongo
	Rpc          RpcServer
	EventStream  EventStream
	Nats         NatsConfig
}
