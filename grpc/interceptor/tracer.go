package interceptor

import (
	"context"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
	"google.golang.org/grpc"

	grpcotel "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc"
)

// Tracer is the interceptor that sets the default values of each input and
// output parameter
var Tracer = &TraceHandler{}

// TraceHandler represents a defaulter
type TraceHandler struct{}

// Unary does unary validation
func (h *TraceHandler) Unary(ctx context.Context, input interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer h.flush()

	var (
		tracer = h.tracer()
		exec   = grpcotel.UnaryServerInterceptor(tracer)
	)

	return exec(ctx, input, info, handler)
}

// Stream does not validate the stream
func (h *TraceHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	defer h.flush()

	var (
		tracer = h.tracer()
		exec   = grpcotel.StreamServerInterceptor(tracer)
	)

	return exec(srv, stream, info, handler)
}

func (h *TraceHandler) tracer() trace.Tracer {
	provider := global.TraceProvider()

	// tracer
	tracer := provider.Tracer(
		"github.com/phogolabs/plex/grpc",
		trace.WithInstrumentationVersion("0.1"),
	)

	return tracer
}

func (h *TraceHandler) flush() {
	type Flusher interface {
		Flush()
	}

	if flusher, ok := global.TraceProvider().(Flusher); ok {
		flusher.Flush()
	}
}
