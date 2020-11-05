package middleware

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/api/global"
)

// Tracer is the interceptor that sets the default values of each input and
// output parameter
func Tracer(next http.Handler) http.Handler {
	type Flusher interface {
		Flush()
	}

	provider := global.TracerProvider()

	fn := func(w http.ResponseWriter, r *http.Request) {
		// wrap the handler
		handler := otelhttp.NewHandler(next, r.URL.Path)
		handler.ServeHTTP(w, r)

		if flusher, ok := provider.(Flusher); ok {
			flusher.Flush()
		}
	}

	return http.HandlerFunc(fn)
}

// type UnaryClientInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error
