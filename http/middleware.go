package http

import (
	"bufio"
	"io"
	"net/http"

	"github.com/go-chi/chi/middleware"
)

const (
	// ContentTypeAll represents '*/*' content-type
	ContentTypeAll = "*/*"
	// ContentTypeForm is the form url encoded
	ContentTypeForm = "application/x-www-form-urlencoded"
	// ContentTypeGRPC represents 'application/grpc' content-type
	ContentTypeGRPC = "application/grpc"
	// ContentTypeJSON represents 'application/json' content-type
	ContentTypeJSON = "application/json"
)

// Forwarder middleware injects some useful headers
func Forwarder(next http.Handler) http.Handler {
	prepare := func(r *http.Request, header string) {
		if value := r.Header.Get(header); value == "" || value == ContentTypeAll {
			r.Header.Set(header, ContentTypeJSON)
		}
	}

	fn := func(w http.ResponseWriter, r *http.Request) {
		// set some metadata headers
		r.Header.Set("X-Plex-Real-Ip", r.RemoteAddr)
		r.Header.Set("X-Plex-Request-Id", middleware.GetReqID(r.Context()))
		r.Header.Set("X-Plex-User-Agent", r.UserAgent())

		// prepare the default header if thery are not set
		prepare(r, "Accept")
		prepare(r, "Content-Type")

		// serve the request
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// Match matches the request
func Match(body io.Reader) bool {
	r, err := http.ReadRequest(bufio.NewReader(body))
	if err != nil {
		return false
	}

	if value := r.Header.Get("Content-Type"); value != ContentTypeGRPC {
		return true
	}

	return false
}
