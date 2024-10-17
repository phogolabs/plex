package interceptor

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

var (
	// ClientUnaryTracer intercepts client unary connections
	ClientUnaryTracer = grpc.WithUnaryInterceptor(
		otelgrpc.UnaryClientInterceptor(
			otelgrpc.WithSpanOptions(
				trace.WithSpanKind(trace.SpanKindInternal),
			),
		))

	// ClientStreamTracer intercepts client stream connections
	ClientStreamTracer = grpc.WithStreamInterceptor(
		otelgrpc.StreamClientInterceptor(
			otelgrpc.WithSpanOptions(
				trace.WithSpanKind(trace.SpanKindInternal),
			),
		))
)

// Tracer is the interceptor that sets the default values of each input and
// output parameter
var Tracer = &TraceHandler{
	Options: []otelgrpc.Option{
		otelgrpc.WithSpanOptions(
			trace.WithSpanKind(trace.SpanKindInternal),
		),
	},
}

// TraceHandler represents a defaulter
type TraceHandler struct {
	Options []otelgrpc.Option
}

// Unary does unary validation
func (h *TraceHandler) Unary(ctx context.Context, input interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	exec := otelgrpc.UnaryServerInterceptor(h.Options...)
	return exec(ctx, input, info, handler)
}

// Stream does not validate the stream
func (h *TraceHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	exec := otelgrpc.StreamServerInterceptor(h.Options...)
	return exec(srv, stream, info, handler)
}
