package interceptor

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

var tracerOptions = []otelgrpc.Option{
	otelgrpc.WithSpanOptions(
		trace.WithSpanKind(trace.SpanKindInternal),
	),
}

// ClientTracer is a dial option that enables tracing for gRPC client connections
var ClientTracer = grpc.WithStatsHandler(otelgrpc.NewClientHandler(tracerOptions...))

// ServerTracer is a server option that enables tracing for gRPC server handlers
var ServerTracer = grpc.StatsHandler(otelgrpc.NewServerHandler(tracerOptions...))
