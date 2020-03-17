package grpc

import (
	"context"

	"google.golang.org/grpc"
)

// Interceptor represents a grpc interface
type Interceptor interface {
	Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	Stream(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
}

// ChainInterceptor creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryServer(one, two, three) will execute one before two before three, and three
// will see context changes of one and two.
type ChainInterceptor []Interceptor

// Unary creates a single unary interceptor out of a chain of many interceptors.
func (ch ChainInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var (
		interceptors = ch
		n            = len(interceptors)
	)

	chainer := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
		return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
			return currentInter(currentCtx, currentReq, info, currentHandler)
		}
	}

	chainedHandler := handler
	for i := n - 1; i >= 0; i-- {
		chainedHandler = chainer(interceptors[i].Unary, chainedHandler)
	}

	return chainedHandler(ctx, req)
}

// Stream creates a single stream interceptor out of a chain of many interceptors.
func (ch ChainInterceptor) Stream(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	var (
		interceptors = ch
		n            = len(interceptors)
	)

	chainer := func(currentInter grpc.StreamServerInterceptor, currentHandler grpc.StreamHandler) grpc.StreamHandler {
		return func(currentSrv interface{}, currentStream grpc.ServerStream) error {
			return currentInter(currentSrv, currentStream, info, currentHandler)
		}
	}

	chainedHandler := handler
	for i := n - 1; i >= 0; i-- {
		chainedHandler = chainer(interceptors[i].Stream, chainedHandler)
	}

	return chainedHandler(srv, ss)
}
