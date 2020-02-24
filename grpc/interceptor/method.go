package interceptor

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"
)

// Method is the method handler
var Method = &MethodHandler{}

var (
	// MethodCtxKey is the context.Context key to store the request log entry.
	MethodCtxKey = &contextKey{"method"}
)

// MethodContext allows a Receiver to understand the context of a request.
type MethodContext struct {
	Handler string
	Method  string
}

// SetMethodContext sets a method entry into the provided context
func SetMethodContext(ctx context.Context, m *MethodContext) context.Context {
	return context.WithValue(ctx, MethodCtxKey, m)
}

// GetMethodContext returns the method Entry found in the context,
// or a new Default log Entry if none is found
func GetMethodContext(ctx context.Context) *MethodContext {
	v := ctx.Value(MethodCtxKey)

	if m, ok := v.(*MethodContext); ok {
		return m
	}

	return nil
}

// MethodHandler enriches the execution context with method context
type MethodHandler struct{}

// Unary does unary logging
func (l *MethodHandler) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	mCtx := &MethodContext{
		Handler: l.name(info.Server),
		Method:  info.FullMethod,
	}

	ctx = SetMethodContext(ctx, mCtx)
	return handler(ctx, req)
}

// Stream does stream logging
func (l *MethodHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := stream.Context()

	mCtx := &MethodContext{
		Handler: l.name(srv),
		Method:  info.FullMethod,
	}

	stream = &ServerStream{
		Ctx:    SetMethodContext(ctx, mCtx),
		Stream: stream,
	}

	return handler(srv, stream)
}

func (l *MethodHandler) name(srv interface{}) string {
	return strings.TrimPrefix(fmt.Sprintf("%T", srv), "*")
}
