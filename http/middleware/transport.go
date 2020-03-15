package middleware

import (
	"net/http"
	"net/textproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ServerTransportStream represents a http stream
type ServerTransportStream struct {
	method string
	header http.Header
}

// Method returns the method
func (s *ServerTransportStream) Method() string {
	return s.method
}

// SetHeader sets the header
func (s *ServerTransportStream) SetHeader(md metadata.MD) error {
	for key, values := range md {
		for _, value := range values {
			s.header.Set(key, value)
		}
	}

	return nil
}

// SendHeader sends the header
func (s *ServerTransportStream) SendHeader(md metadata.MD) error {
	for key, values := range md {
		for _, value := range values {
			s.header.Add(key, value)
		}
	}

	return nil
}

// SetTrailer sets the trailer header
func (s *ServerTransportStream) SetTrailer(md metadata.MD) error {
	for key := range md {
		prefixed := textproto.CanonicalMIMEHeaderKey(key)
		s.header.Add("Trailer", prefixed)
	}

	return nil
}

// ServerTransport injects the server transport stream
func ServerTransport(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var (
			ts = &ServerTransportStream{
				method: r.URL.Path,
				header: w.Header(),
			}

			ctx = grpc.NewContextWithServerTransportStream(r.Context(), ts)
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
