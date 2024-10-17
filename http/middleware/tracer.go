package middleware

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// Tracer is the interceptor that sets the default values of each input and
// output parameter
func Tracer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		options := []otelhttp.Option{
			otelhttp.WithServerName(r.Host),
		}

		// prepare the handler
		handler := otelhttp.NewHandler(next, r.Host, options...)
		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
