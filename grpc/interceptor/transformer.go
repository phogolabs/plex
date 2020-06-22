package interceptor

import (
	"context"

	"github.com/go-playground/mold"
	"google.golang.org/grpc"
)

// Transformer is the interceptor that transforms the values of each input and
// output parameter
var Transformer = &TransformHandler{
	Molder: mold.New(),
}

// TransformHandler represents a mold
type TransformHandler struct {
	Molder *mold.Transformer
}

// Unary does unary validation
func (h *TransformHandler) Unary(ctx context.Context, input interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// transform input
	if input != nil {
		if err := h.Molder.Struct(ctx, input); err != nil {
			return nil, err
		}
	}

	// execute handler
	output, err := handler(ctx, input)

	// transform output
	if output != nil {
		if err := h.Molder.Struct(ctx, output); err != nil {
			return nil, err
		}
	}

	return output, err
}

// Stream does not validate the stream
func (h *TransformHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, stream)
}
