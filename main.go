package main

import (
	"net/http"

	"example.com/fxdemo/handlers"
	"example.com/fxdemo/server"
	"example.com/fxdemo/wire"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			server.NewHTTPServer,
			fx.Annotate(
				server.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			wire.AsRoute(handlers.NewHelloHandler),
			wire.AsRoute(handlers.NewEchoHandler),
			zap.NewExample,
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.With(zap.String("logger", "fx-example"))
		}),
		fx.Invoke(func(srv *http.Server) {}),
	).Run()
}
