package grpc

import (
	"context"

	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

// Gateway represents a grpc server
type Gateway struct {
	Server  *grpc.Server
	Options []grpc.ServerOption
}

// NewGateway creates a new grpc server
func NewGateway(opts ...GatewayOption) *Gateway {
	if len(opts) == 0 {
		opts = append(opts, WithDefault())
	}

	gateway := &Gateway{
		Options: []grpc.ServerOption{
			KeepaliveOption,
		},
	}

	for _, op := range opts {
		op.Apply(gateway)
	}

	gateway.Server = grpc.NewServer(gateway.Options...)
	return gateway
}

// Serve serves the mux
func (gateway *Gateway) Serve(mux cmux.CMux) error {
	listener := mux.MatchWithWriters(Match)
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
