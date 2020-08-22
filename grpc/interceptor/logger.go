package interceptor

import (
	"context"
	"strings"
	"time"

	"github.com/phogolabs/log"
	"go.opentelemetry.io/otel/api/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Logger is the log interceptor
var Logger = &LogHandler{}

// LogHandler represents a logger
type LogHandler struct{}

// Unary does unary logging
func (h *LogHandler) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fields := annotation(ctx)
	fields["method"] = info.FullMethod

	logger := log.GetContext(ctx)
	logger = logger.WithFields(fields)

	start := time.Now()
	ctx = log.SetContext(ctx, logger)
	response, err := handler(ctx, req)

	logger = logger.WithFields(log.Map{
		"duration": time.Since(start),
	})

	if err != nil {
		logger.WithError(err).Error("executing method fail")
		return nil, err
	}

	logger.Info("executing method success")
	return response, nil

}

// Stream does stream logging
func (h *LogHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	var (
		ctx    = stream.Context()
		logger = log.GetContext(ctx)
		source = "server"
	)

	if info.IsClientStream {
		source = "client"
	}

	fields := annotation(ctx)
	fields["source"] = source
	fields["method"] = info.FullMethod

	logger = logger.WithFields(fields)

	stream = &ServerStream{
		Ctx:    log.SetContext(ctx, logger),
		Stream: stream,
	}

	start := time.Now()
	err := handler(srv, stream)

	logger = logger.WithFields(log.Map{
		"duration": time.Since(start),
	})

	if err != nil {
		logger.WithError(err).Error("streaming method fail")
		return err
	}

	logger.Info("streaming method success")
	return nil
}

func annotation(ctx context.Context) log.Map {
	fields := log.Map{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		for k, v := range md {
			if strings.HasPrefix(k, "x-plex") && len(v) > 0 {
				fields[k] = v[0]
			}
		}
	}

	span := trace.
		SpanFromContext(ctx).
		SpanContext()

	if traceID := span.TraceID; len(traceID) > 0 {
		fields["trace_id"] = traceID
	}

	if spanID := span.SpanID; len(spanID) > 0 {
		fields["span_id"] = spanID
	}

	return fields
}
