package grpc

import (
	"context"

	"github.com/phogolabs/plex/grpc/interceptor"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

const (
	// ContentType is a grpc content-type
	ContentType = "application/grpc"
)

// Gateway represents a grpc server
type Gateway struct {
	Server *grpc.Server
}

// NewGateway creates a new grpc server
func NewGateway() *Gateway {
	chain := ChainInterceptor{
		interceptor.Logger,
		interceptor.Recoverer,
		interceptor.Validator,
	}

	return &Gateway{
		Server: grpc.NewServer(
			grpc.UnaryInterceptor(chain.Unary),
			grpc.StreamInterceptor(chain.Stream),
		),
	}
}

// Serve serves the mux
func (gateway *Gateway) Serve(mux cmux.CMux) error {
	// listener := mux.Match(cmux.HTTP2HeaderField("content-type", ContentType))

	listener := mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", ContentType))
	return gateway.Server.Serve(listener)
}

// Shutdown shutdowns the server
func (gateway *Gateway) Shutdown(ctx context.Context) error {
	if gateway.Server == nil {
		return nil
	}

	gateway.Server.GracefulStop()
	return nil
}
