package plex

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// SetCookie adds a Set-Cookie header to the provided server transport context.
// The provided cookie must have a valid Name. Invalid cookies may be
// silently dropped.
func SetCookie(ctx context.Context, cookie *http.Cookie) {
	if v := cookie.String(); v != "" {
		md := metadata.Pairs("Set-Cookie", v)
		grpc.SetHeader(ctx, md)
	}
}

// CookieFromIncomingContext returns the cookie
func CookieFromIncomingContext(ctx context.Context, name string) (*http.Cookie, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, http.ErrNoCookie
	}

	value, ok := md["cookie"]
	if !ok {
		return nil, http.ErrNoCookie
	}

	r := &http.Request{
		Header: http.Header{
			"Cookie": value,
		},
	}

	return r.Cookie(name)
}
