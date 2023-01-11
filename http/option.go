package http

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/proto"
)

type (
	// HeaderMatcherFunc checks whether a header key should be forwarded to/from gRPC context.
	HeaderMatcherFunc = runtime.HeaderMatcherFunc

	// ErrorHandlerFunc handles the error as a gRPC error generated via status package and replies to the request.
	ErrorHandlerFunc = runtime.ErrorHandlerFunc
)

var (
	// WithForwardResponseOption returns a ServeMuxOption representing the forwardResponseOption.
	//
	// forwardResponseOption is an option that will be called on the relevant context.Context,
	// http.ResponseWriter, and proto.Message before every forwarded response.
	//
	// The message may be nil in the case where just a header is being sent.
	WithForwardResponseOption = runtime.WithForwardResponseOption

	// WithIncomingHeaderMatcher returns a ServeMuxOption representing a headerMatcher for incoming request to gateway.
	//
	// This matcher will be called with each header in http.Request. If matcher returns true, that header will be
	// passed to gRPC context. To transform the header before passing to gRPC context, matcher should return modified header.
	WithIncomingHeaderMatcher = runtime.WithIncomingHeaderMatcher

	// WithOutgoingHeaderMatcher returns a ServeMuxOption representing a headerMatcher for outgoing response from gateway.
	//
	// This matcher will be called with each header in response header metadata. If matcher returns true, that header will be
	// passed to http response returned from gateway. To transform the header before passing to response,
	// matcher should return modified header.
	WithOutgoingHeaderMatcher = runtime.WithOutgoingHeaderMatcher

	// WithMetadata returns a ServeMuxOption for passing metadata to a gRPC context.
	//
	// This can be used by services that need to read from http.Request and modify gRPC context. A common use case
	// is reading token from cookie and adding it in gRPC context.
	WithMetadata = runtime.WithMetadata

	// WithErrorHandler returns a ServeMuxOption for passing metadata to a gRPC context.
	//
	// This can be used to handle an error as general proto message defined by gRPC.
	// The response including body and status is not backward compatible with the default error handler.
	// When this option is used, HTTPError and OtherErrorHandler are overwritten on initialization.
	WithErrorHandler = runtime.WithErrorHandler

	// WithDisablePathLengthFallback returns a ServeMuxOption for disable path length fallback.
	WithDisablePathLengthFallback = runtime.WithDisablePathLengthFallback

	// WithStreamErrorHandler returns a ServeMuxOption that will use the given custom stream
	// error handler, which allows for customizing the error trailer for server-streaming
	// calls.
	//
	// For stream errors that occur before any response has been written, the mux's
	// ProtoErrorHandler will be invoked. However, once data has been written, the errors must
	// be handled differently: they must be included in the response body. The response body's
	// final message will include the error details returned by the stream error handler.
	WithStreamErrorHandler = runtime.WithStreamErrorHandler

	// WithMarshaler returns a ServeMuxOption which associates inbound and outbound
	// Marshalers to a MIME type in mux.
	WithMarshaler = runtime.WithMarshalerOption
)

var (
	// AllIncomingHeaders allows the service to handle all incoming request
	AllIncomingHeaders = runtime.WithIncomingHeaderMatcher(preserve)

	// AllOutgoingHeaders allows the service to handle all incoming request
	AllOutgoingHeaders = runtime.WithOutgoingHeaderMatcher(preserve)

	// WithForwardResponse represents a forward response
	WithForwardResponse = runtime.WithForwardResponseOption(response)
)

func preserve(value string) (string, bool) {
	if strings.EqualFold(value, "connection") {
		return "", false
	}

	return value, true
}

func response(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	kv, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	// set http status code
	if value := kv.HeaderMD.Get("x-http-code"); len(value) > 0 {
		var (
			header   = w.Header()
			headerKV = kv.HeaderMD
		)
		code, err := strconv.Atoi(value[0])
		if err != nil {
			return err
		}

		delete(header, "Grpc-Metadata-X-Http-Code")
		delete(headerKV, "x-http-code")

		w.WriteHeader(code)
	}

	return nil
}
