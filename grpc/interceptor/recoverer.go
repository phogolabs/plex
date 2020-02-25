package interceptor

import (
	"context"

	"github.com/phogolabs/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Recover is the recovery interceptor
var Recover = &RecoveryHandler{}

// RecoveryHandlerFuncContext is a function that recovers from the panic `p` by returning an `error`.
// The context can be used to extract request scoped metadata and context values.
type RecoveryHandlerFuncContext func(ctx context.Context, p interface{}) (err error)

// RecoveryHandler represents an interceptor that recovers from panic
type RecoveryHandler struct {
	Handler RecoveryHandlerFuncContext
}

// Unary does unary logging
func (l *RecoveryHandler) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
	logger := log.GetContext(ctx)

	defer func() {
		if r := recover(); r != nil {
			if err = l.recoverFrom(ctx, r); err != nil {
				logger.WithError(err).Error("fatal error ocurred")
			}
		}
	}()

	return handler(ctx, req)
}

// Stream does stream logging
func (l *RecoveryHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	var (
		ctx    = stream.Context()
		logger = log.GetContext(ctx)
	)

	defer func() {
		if r := recover(); r != nil {
			if err = l.recoverFrom(ctx, r); err != nil {
				logger.WithError(err).Error("fatal error ocurred")
			}
		}
	}()

	return handler(srv, stream)
}

func (l *RecoveryHandler) recoverFrom(ctx context.Context, p interface{}) error {
	if l.Handler == nil {
		return status.Errorf(codes.Internal, "%s", p)
	}

	return l.Handler(ctx, p)
}