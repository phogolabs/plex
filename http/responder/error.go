package responder

import (
	"context"
	"net/http"

	"github.com/goware/errorx"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PostgreSQLFormatter can be used to handle an error as general proto message defined by gRPC.
func PostgreSQLFormatter(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	statusErr, ok := status.FromError(err)
	if !ok {
		statusErr = status.New(codes.Unknown, err.Error())
	}

	code := runtime.HTTPStatusFromCode(statusErr.Code())
	w.WriteHeader(code)

	errx := errorx.New(code, http.StatusText(code), statusErr.Message())
	marshaler.NewEncoder(w).Encode(errx)
}
