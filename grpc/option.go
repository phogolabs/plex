package grpc

import (
	"time"

	"github.com/phogolabs/plex/grpc/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// KeepaliveOption represents a keep alive option
// For more information have a look:
//
//	https://stackoverflow.com/questions/52993259/problem-with-grpc-setup-getting-an-intermittent-rpc-unavailable-error
var KeepaliveOption = grpc.KeepaliveParams(keepalive.ServerParameters{
	MaxConnectionIdle: 5 * time.Minute,
})

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
		gw.Options = append(gw.Options, grpc.UnaryInterceptor(chain.Unary))
		gw.Options = append(gw.Options, grpc.StreamInterceptor(chain.Stream))
	}

	return GatewayOptionFunc(fn)
}
