package middleware

import (
	"net/http"

	httpie "go.opentelemetry.io/contrib/instrumentation/net/http"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
)

// Tracer is the interceptor that sets the default values of each input and
// output parameter
func Tracer(next http.Handler) http.Handler {
	type Flusher interface {
		Flush()
	}

	provider := global.TraceProvider()

	// tracer
	tracer := provider.Tracer(
		"github.com/phogolabs/plex/http",
		trace.WithInstrumentationVersion("0.1"))

	fn := func(w http.ResponseWriter, r *http.Request) {
		// wrap the handler
		handler := httpie.NewHandler(
			next, r.URL.Path,
			httpie.WithTracer(tracer),
		)

		handler.ServeHTTP(w, r)

		if flusher, ok := provider.(Flusher); ok {
			flusher.Flush()
		}
	}

	return http.HandlerFunc(fn)
}
