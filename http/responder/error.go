package responder

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/inflect"
	"github.com/goware/errorx"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

// PostgreSQLErrorFormatter can be used to handle an error as general proto message defined by gRPC.
func PostgreSQLErrorFormatter(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	statusErr, ok := status.FromError(err)
	if !ok {
		statusErr = status.New(codes.Unknown, err.Error())
	}

	code := runtime.HTTPStatusFromCode(statusErr.Code())
	errx := errorx.New(code, statusErr.Message())

	// add details
	for _, d := range statusErr.Details() {
		if obj, ok := d.(*structpb.Struct); ok {
			for _, v := range obj.AsMap() {
				errx.Details = append(errx.Details, fmt.Sprintf("%v", v))
			}
		}
	}

	// set the response status code
	w.WriteHeader(code)

	// set the response body
	marshaler.NewEncoder(w).Encode(errx)
}

// CanonicalErrorFormatter can be used to handle an error as general proto message defined by gRPC.
func CanonicalErrorFormatter(_ context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	errx, _ := status.FromError(err)

	type Response struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	response := &Response{
		Code:    errx.Code().String(),
		Message: errx.Message(),
	}

	response.Code = inflect.Underscore(response.Code)
	response.Code = strings.ToUpper(response.Code)

	// set the response status code
	w.WriteHeader(runtime.HTTPStatusFromCode(errx.Code()))

	// set the response body
	marshaler.NewEncoder(w).Encode(response)
}
