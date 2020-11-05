package interceptor

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/api/global"
	"google.golang.org/grpc"
)

// Tracer is the interceptor that sets the default values of each input and
// output parameter
var Tracer = &TraceHandler{}

// TraceHandler represents a defaulter
type TraceHandler struct{}

// Unary does unary validation
func (h *TraceHandler) Unary(ctx context.Context, input interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer h.flush()

	exec := otelgrpc.UnaryServerInterceptor()
	return exec(ctx, input, info, handler)
}

// Stream does not validate the stream
func (h *TraceHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	defer h.flush()

	exec := otelgrpc.StreamServerInterceptor()
	return exec(srv, stream, info, handler)
}

func (h *TraceHandler) flush() {
	type Flusher interface {
		Flush()
	}

	if flusher, ok := global.TracerProvider().(Flusher); ok {
		flusher.Flush()
	}
}
