package interceptor

import (
	"context"
	"reflect"

	"github.com/phogolabs/inflate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Defaulter is the interceptor that sets the default values of each input and
// output parameter
var Defaulter = &DefaultHandler{}

// DefaultHandler represents a defaulter
type DefaultHandler struct{}

// Unary does unary validation
func (h *DefaultHandler) Unary(ctx context.Context, input interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if h.canSet(input) {
		if err := inflate.SetDefault(input); err != nil {
			err = status.Error(codes.Internal, err.Error())
			return nil, err
		}
	}

	output, err := handler(ctx, input)

	if h.canSet(output) {
		if err := inflate.SetDefault(output); err != nil {
			err = status.Error(codes.Internal, err.Error())
			return nil, err
		}
	}

	return output, err
}

// Stream does not validate the stream
func (h *DefaultHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, stream)
}

func (h *DefaultHandler) canSet(input interface{}) bool {
	value := reflect.ValueOf(input)
	return value.Kind() == reflect.Ptr && !value.IsNil()
}
