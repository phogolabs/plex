package meta

import "context"

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

// NewTokenContext sets the session token to a given context
func NewTokenContext(ctx context.Context, token *Token) context.Context {
	ctx = context.WithValue(ctx, TokenCtxKey, token)
	return ctx
}

// TokenFromContext returns a session token
func TokenFromContext(ctx context.Context) *Token {
	if token, ok := ctx.Value(TokenCtxKey).(*Token); ok {
		return token
	}

	return nil
}
