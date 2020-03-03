package plex

import "github.com/phogolabs/plex/http"

// ServerOption represents a server option
type ServerOption interface {
	// Apply applies the option
	Apply(*Server)
}

// ServerOptionFunc represents a server option as function
type ServerOptionFunc func(*Server)

// Apply applies the option
func (op ServerOptionFunc) Apply(srv *Server) {
	op(srv)
}

// WithAddress sets the server address
func WithAddress(addr string) ServerOption {
	fn := func(srv *Server) {
		srv.addr = addr
	}

	return ServerOptionFunc(fn)
}

type (
	// HeaderMatcherFunc checks whether a header key should be forwarded to/from gRPC context.
	HeaderMatcherFunc = http.HeaderMatcherFunc

	// ProtoErrorHandlerFunc handles the error as a gRPC error generated via status package and replies to the request.
	ProtoErrorHandlerFunc = http.ProtoErrorHandlerFunc
)

var (
	// WithForwardResponseOption returns a ServeMuxOption representing the forwardResponseOption.
	//
	// forwardResponseOption is an option that will be called on the relevant context.Context,
	// http.ResponseWriter, and proto.Message before every forwarded response.
	//
	// The message may be nil in the case where just a header is being sent.
	WithForwardResponseOption = http.WithForwardResponseOption

	// WithIncomingHeaderMatcher returns a ServeMuxOption representing a headerMatcher for incoming request to gateway.
	//
	// This matcher will be called with each header in http.Request. If matcher returns true, that header will be
	// passed to gRPC context. To transform the header before passing to gRPC context, matcher should return modified header.
	WithIncomingHeaderMatcher = http.WithIncomingHeaderMatcher

	// WithOutgoingHeaderMatcher returns a ServeMuxOption representing a headerMatcher for outgoing response from gateway.
	//
	// This matcher will be called with each header in response header metadata. If matcher returns true, that header will be
	// passed to http response returned from gateway. To transform the header before passing to response,
	// matcher should return modified header.
	WithOutgoingHeaderMatcher = http.WithOutgoingHeaderMatcher

	// WithMetadata returns a ServeMuxOption for passing metadata to a gRPC context.
	//
	// This can be used by services that need to read from http.Request and modify gRPC context. A common use case
	// is reading token from cookie and adding it in gRPC context.
	WithMetadata = http.WithMetadata

	// WithProtoErrorHandler returns a ServeMuxOption for passing metadata to a gRPC context.
	//
	// This can be used to handle an error as general proto message defined by gRPC.
	// The response including body and status is not backward compatible with the default error handler.
	// When this option is used, HTTPError and OtherErrorHandler are overwritten on initialization.
	WithProtoErrorHandler = http.WithProtoErrorHandler

	// WithDisablePathLengthFallback returns a ServeMuxOption for disable path length fallback.
	WithDisablePathLengthFallback = http.WithDisablePathLengthFallback

	// WithStreamErrorHandler returns a ServeMuxOption that will use the given custom stream
	// error handler, which allows for customizing the error trailer for server-streaming
	// calls.
	//
	// For stream errors that occur before any response has been written, the mux's
	// ProtoErrorHandler will be invoked. However, once data has been written, the errors must
	// be handled differently: they must be included in the response body. The response body's
	// final message will include the error details returned by the stream error handler.
	WithStreamErrorHandler = http.WithStreamErrorHandler

	// WithLastMatchWins returns a ServeMuxOption that will enable "last
	// match wins" behavior, where if multiple path patterns match a
	// request path, the last one defined in the .proto file will be used.
	WithLastMatchWins = http.WithLastMatchWins
)

// WithServeMuxOption creates a server option with mux options
func WithServeMuxOption(opts ...http.ServeMuxOption) ServerOption {
	fn := func(srv *Server) {
		srv.httpSrv = http.NewServer(opts...)
	}

	return ServerOptionFunc(fn)
}
