package interceptor

import (
	"context"

	"github.com/go-playground/mold"
	"google.golang.org/grpc"
)

// Transformer is the interceptor that transforms the values of each input and
// output parameter
var Transformer = &TransformHandler{
	Transformer: mold.New(),
}

// TransformHandler represents a mold
type TransformHandler struct {
	Transformer *mold.Transformer
}

// Register adds a transformation with the given tag
func (h *TransformHandler) Register(tag string, fn mold.Func) {
	h.Transformer.Register(tag, fn)
}

// RegisterStruct registers a StructLevelFunc against a number of types.
// Why does this exist? For structs for which you may not have access or rights to add tags too,
// from other packages your using.
func (h *TransformHandler) RegisterStruct(fn mold.StructLevelFunc, types ...interface{}) {
	h.Transformer.RegisterStructLevel(fn, types...)
}

// Unary does unary validation
func (h *TransformHandler) Unary(ctx context.Context, input interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// transform input
	if input != nil {
		if err := h.Transformer.Struct(ctx, input); err != nil {
			return nil, err
		}
	}

	// execute handler
	output, err := handler(ctx, input)

	// transform output
	if output != nil {
		if err := h.Transformer.Struct(ctx, output); err != nil {
			return nil, err
		}
	}

	return output, err
}

// Stream does not validate the stream
func (h *TransformHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, stream)
}
