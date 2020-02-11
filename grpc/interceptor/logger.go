package interceptor

import (
	"context"
	"fmt"
	"strings"

	"github.com/phogolabs/log"
	"google.golang.org/grpc"
)

// Logger is the log interceptor
var Logger = &LogHandler{}

// LogHandler represents a logger
type LogHandler struct{}

// Unary does unary logging
func (l *LogHandler) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logger := log.GetContext(ctx)

	logger = logger.WithFields(log.Map{
		"handler": l.name(info.Server),
		"method":  info.FullMethod,
	})

	logger.Info("executing method")
	response, err := handler(ctx, req)

	if err != nil {
		logger.WithError(err).Error("executing method fail")
		return nil, err
	}

	logger.Info("executing method finish")
	return response, nil

}

// Stream does stream logging
func (l *LogHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	var (
		ctx    = stream.Context()
		logger = log.GetContext(ctx)
		source = "server"
	)

	if info.IsClientStream {
		source = "client"
	}

	logger = logger.WithFields(log.Map{
		"source":  source,
		"handler": l.name(srv),
		"method":  info.FullMethod,
	})

	logger.Info("streaming method")
	err := handler(srv, stream)

	if err != nil {
		logger.WithError(err).Error("streaming fail")
		return err
	}

	logger.Info("streaming finish")
	return nil
}

func (l *LogHandler) name(srv interface{}) string {
	return strings.TrimPrefix(fmt.Sprintf("%T", srv), "*")
}
