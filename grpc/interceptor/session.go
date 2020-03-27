package interceptor

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// TokenCtxKey represents the session ctx
var TokenCtxKey = &contextKey{name: "token"}

// Token represents the session token
type Token struct {
	AuthTime int64                  `json:"auth_time"`
	Issuer   string                 `json:"iss"`
	Audience string                 `json:"aud"`
	Expires  int64                  `json:"exp"`
	IssuedAt int64                  `json:"iat"`
	Subject  string                 `json:"sub,omitempty"`
	UID      string                 `json:"uid,omitempty"`
	Claims   map[string]interface{} `json:"claims"`
}

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

	ctx = SetToken(ctx, token)
	return handler(ctx, req)
}

// Stream does stream logging
func (h *SessionHandler) Stream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := stream.Context()

	token, err := h.session(ctx)
	if err != nil {
		return err
	}

	ctx = SetToken(ctx, token)

	stream = &ServerStream{
		Ctx:    ctx,
		Stream: stream,
	}

	return handler(srv, stream)
}

func (h *SessionHandler) session(ctx context.Context) (*Token, error) {
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

	token := &Token{}
	if err := decoder.Decode(token); err != nil {
		return nil, unauth
	}

	return token, nil
}

// SetToken sets the session token to a given context
func SetToken(ctx context.Context, token *Token) context.Context {
	ctx = context.WithValue(ctx, TokenCtxKey, token)
	return ctx
}

// GetToken returns a session token
func GetToken(ctx context.Context) *Token {
	if token, ok := ctx.Value(TokenCtxKey).(*Token); ok {
		return token
	}

	return nil
}
