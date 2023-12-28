package wire

import (
	"example.com/fxdemo/server"
	"go.uber.org/fx"
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(server.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
