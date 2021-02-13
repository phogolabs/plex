package interceptor

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

var (
	// ClientUnaryTracer intercepts client unary connections
	ClientUnaryTracer = grpc.WithUnaryInterceptor(
		otelgrpc.UnaryClientInterceptor())

	// ClientStreamTracer intercepts client stream connections
	ClientStreamTracer = grpc.WithStreamInterceptor(
		otelgrpc.StreamClientInterceptor())
)

// Tracer is the interceptor that sets the default values of each input and
// output parameter
var Tracer = &TraceHandler{}

// TraceHandler represents a defaulter
type TraceHandler struct{}

// Unary does unary validation
func (h *TraceHandler) Unary(ctx context.Context, input interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	exec := otelgrpc.UnaryServerInterceptor()
	return exec(ctx, input, info, handler)
}

// Stream does not validate the stream
func (h *TraceHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	exec := otelgrpc.StreamServerInterceptor()
	return exec(srv, stream, info, handler)
}
