package interceptor

import (
	"context"

	validate "github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Validator is the validation interceptor
var Validator = &ValidationHandler{}

// ValidationHandler represents a logger
type ValidationHandler struct{}

// Unary does unary validation
func (l *ValidationHandler) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := validate.New().StructCtx(ctx, req); err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}

	return handler(ctx, req)
}

// Stream does not validate the stream
func (l *ValidationHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, stream)
}
