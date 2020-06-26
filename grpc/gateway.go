package grpc

import (
	"context"
	"time"

	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	// ContentType is a grpc content-type
	ContentType = "application/grpc"
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
			// For more information have a look:
			//   https://stackoverflow.com/questions/52993259/problem-with-grpc-setup-getting-an-intermittent-rpc-unavailable-error
			grpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionIdle: 5 * time.Minute,
			}),
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
	listener := mux.MatchWithWriters(
		cmux.HTTP2MatchHeaderFieldSendSettings("content-type", ContentType),
	)

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
