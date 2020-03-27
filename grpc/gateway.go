package grpc

import (
	"context"

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
func NewGateway(opts ...GatewayOption) *Gateway {
	if len(opts) == 0 {
		opts = append(opts, WithDefault())
	}

	gateway := &Gateway{
		Server: grpc.NewServer(),
	}

	for _, op := range opts {
		op.Apply(gateway)
	}

	return gateway
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
