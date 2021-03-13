package http

import (
	"bufio"
	"io"
	"mime"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/phogolabs/log"
)

var (
	// HeaderContentType represents a header Content-Type key
	HeaderContentType = textproto.CanonicalMIMEHeaderKey("Content-Type")

	// HeaderAccept represents a header Accept key
	HeaderAccept = textproto.CanonicalMIMEHeaderKey("Accept")
)

const (
	// ContentTypeAll is the '*/*' type
	ContentTypeAll = "*/*"
	// ContentTypeForm is the form url encoded
	ContentTypeForm = "application/x-www-form-urlencoded"
	// ContentTypeGRPC represents 'application/grpc' content-type
	ContentTypeGRPC = "application/grpc"
	// ContentTypeJSON represents 'application/json' content-type
	ContentTypeJSON = "application/json"
)

// PrepareMediaType prepares a media type header
func PrepareMediaType(name string, r *http.Request) {
	const separator = ","

	var (
		logger = log.GetContext(r.Context())
		header = http.Header{}
	)

	// parse the header
	for _, content := range r.Header[name] {
		// skip empty entries
		if len(content) == 0 {
			continue
		}

		for _, item := range strings.Split(content, separator) {
			// skip empty entries
			if len(item) == 0 {
				continue
			}

			// parse the media type
			value, _, err := mime.ParseMediaType(item)
			if err != nil {
				logger.WithError(err).Infof("skip unsupported media type '%v'", item)
				continue
			}

			// skip the all header type because we will override it anyway
			if strings.EqualFold(value, ContentTypeAll) {
				continue
			}

			header.Add(name, value)
		}
	}

	// delete the header
	r.Header.Del(name)

	// set the new header
	if value, ok := header[name]; ok {
		r.Header[name] = value
	}
}

// SetMediaType sets the media type
func SetMediaType(name string, r *http.Request) {
	value, ok := r.Header[name]

	if !ok || len(value) == 0 {
		r.Header.Set(name, ContentTypeJSON)
	}
}

// Accept prepare the Accept header for underlying requests
func Accept(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// prepare the header values
		PrepareMediaType(HeaderAccept, r)
		// set the header media type
		SetMediaType(HeaderAccept, r)
		// serve the request
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// ContentType prepare the Content-Type header for underlying requests
func ContentType(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// set the header media type
		SetMediaType(HeaderContentType, r)
		// serve the request
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// Metadata middleware injects some useful headers
func Metadata(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// set some metadata headers
		r.Header.Set("X-Plex-Real-Ip", r.RemoteAddr)
		r.Header.Set("X-Plex-Request-Id", middleware.GetReqID(r.Context()))
		r.Header.Set("X-Plex-User-Agent", r.UserAgent())

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

	for _, value := range r.Header[HeaderContentType] {
		if strings.Contains(value, ContentTypeGRPC) {
			return false
		}
	}

	return true
}
