package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Route interface {
	http.Handler

	// Pattern reports the path at which this is registered.
	Pattern() string
}

func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping HTTP server")
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
