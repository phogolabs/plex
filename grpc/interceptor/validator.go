package interceptor

import (
	"context"

	validate "github.com/go-playground/validator/v10"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Validator is the validation interceptor
var Validator = &ValidationHandler{}

// ValidationHandler represents a logger
type ValidationHandler struct {
	Validator *validate.Validate
}

// Unary does unary validation
func (h *ValidationHandler) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	type Validator interface {
		Validate() error
	}

	var err error

	if validator, ok := req.(Validator); ok {
		err = validator.Validate()
	} else {
		err = h.validator().StructCtx(ctx, req)
	}

	if err == nil {
		return handler(ctx, req)
	}

	state, ok := status.FromError(err)
	if !ok {
		state = status.New(codes.InvalidArgument, err.Error())
	}

	switch errx := err.(type) {
	case validate.ValidationErrors:
		badRequest := &errdetails.BadRequest{}

		for _, errf := range errx {
			if errd, ok := errf.(error); ok {
				badRequest.FieldViolations = append(badRequest.FieldViolations,
					&errdetails.BadRequest_FieldViolation{
						Field:       errf.Field(),
						Description: errd.Error(),
					},
				)
			}
		}

		state, _ = state.WithDetails(badRequest)

	}

	return nil, state.Err()
}

// Stream does not validate the stream
func (h *ValidationHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, stream)
}

func (h *ValidationHandler) validator() *validate.Validate {
	if h.Validator != nil {
		return h.Validator
	}

	return validate.New()
}
