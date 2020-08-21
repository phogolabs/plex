package service

import (
	"github.com/phogolabs/plex"
	"github.com/phogolabs/plex/grpc"
	"github.com/phogolabs/plex/grpc/interceptor"
	"github.com/phogolabs/plex/http/middleware"
	"github.com/phogolabs/plex/http/responder"
)

// New creates a new server
func New() *plex.Server {
	// create the plex server
	server := plex.NewServer()
	server.Proxy.OnError(responder.PostgreSQLErrorFormatter)
	server.Proxy.Router().Use(middleware.Tracer)

	server.Gateway = grpc.NewGateway(
		grpc.WithChain(
			grpc.ChainInterceptor{
				interceptor.Tracer,
				interceptor.Logger,
				interceptor.Recoverer,
				interceptor.Defaulter,
				interceptor.Transformer,
				interceptor.Validator,
			},
		),
	)

	// handler setup
	handler := &FooHandler{}
	handler.Mount(server)

	// return the server
	return server
}
