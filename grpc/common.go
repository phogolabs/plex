package grpc

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ToString convert to a string pointer from wrapper
func ToString(value *wrapperspb.StringValue) *string {
	if value != nil {
		return &value.Value
	}

	return nil
}

// GetString converts to a string wrapper from a string pointer
func GetString(value *string) *wrapperspb.StringValue {
	if value != nil {
		return &wrapperspb.StringValue{Value: *value}
	}

	return nil
}

// SetCookie adds a Set-Cookie header to the provided server transport context.
// The provided cookie must have a valid Name. Invalid cookies may be
// silently dropped.
func SetCookie(ctx context.Context, cookie *http.Cookie) error {
	if v := cookie.String(); v != "" {
		md := metadata.Pairs("Set-Cookie", v)
		return grpc.SetHeader(ctx, md)
	}

	return nil
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
