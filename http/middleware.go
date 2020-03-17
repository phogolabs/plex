package http

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

// Forwarder middleware injects some useful headers
func Forwarder(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("X-Plex-Real-Ip", r.RemoteAddr)
		r.Header.Set("X-Plex-Request-Id", middleware.GetReqID(r.Context()))
		r.Header.Set("X-Plex-User-Agent", r.UserAgent())

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
