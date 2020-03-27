package interceptor

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/phogolabs/plex/grpc/meta"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Session injects the session into the request interceptor
var Session = &SessionHandler{}

// SessionHandler represents an interceptor that recovers from panic
type SessionHandler struct{}

// Unary does unary logging
func (h *SessionHandler) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	token, err := h.session(ctx)
	if err != nil {
		return nil, err
	}

	ctx = meta.NewTokenContext(ctx, token)
	return handler(ctx, req)
}

// Stream does stream logging
func (h *SessionHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := stream.Context()

	token, err := h.session(ctx)
	if err != nil {
		return err
	}

	stream = &ServerStream{
		Ctx:    meta.NewTokenContext(ctx, token),
		Stream: stream,
	}

	return handler(srv, stream)
}

func (h *SessionHandler) session(ctx context.Context) (*meta.Token, error) {
	unauth := status.Error(codes.Unauthenticated, "session is not established")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, unauth
	}

	session, ok := md["x-session"]
	if !ok || len(session) == 0 {
		return nil, unauth
	}

	decoder := json.NewDecoder(
		base64.NewDecoder(
			base64.StdEncoding,
			bytes.NewBufferString(session[0]),
		),
	)

	token := &meta.Token{}
	if err := decoder.Decode(token); err != nil {
		return nil, unauth
	}

	return token, nil
}
