package grpc

import (
	"github.com/phogolabs/plex/grpc/interceptor"
	"google.golang.org/grpc"
)

// GatewayOption represents a gateway option
type GatewayOption interface {
	Apply(*Gateway)
}

// GatewayOptionFunc represnts a gateway option func
type GatewayOptionFunc func(*Gateway)

// Apply the option
func (fn GatewayOptionFunc) Apply(gw *Gateway) {
	fn(gw)
}

// WithDefault sets the default option
func WithDefault() GatewayOption {
	chain := ChainInterceptor{
		interceptor.Logger,
		interceptor.Recoverer,
		interceptor.Defaulter,
		interceptor.Transformer,
		interceptor.Validator,
	}

	fn := func(gw *Gateway) {
		interceptor := WithChain(chain)
		interceptor.Apply(gw)
	}

	return GatewayOptionFunc(fn)
}

// WithChain  sets the interceptor chain
func WithChain(chain ChainInterceptor) GatewayOption {
	fn := func(gw *Gateway) {
		gw.Server = grpc.NewServer(
			grpc.UnaryInterceptor(chain.Unary),
			grpc.StreamInterceptor(chain.Stream),
		)
	}

	return GatewayOptionFunc(fn)
}
