package meta

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

// NewIncomingContext creates a new context with incoming md attached.
func NewIncoming(ctx context.Context, kv map[string]string) context.Context {
	return metadata.NewIncomingContext(ctx, metadata.New(kv))
}

// NewOutgoingContext creates a new context with outgoing md attached. If used
// in conjunction with AppendToOutgoingContext, NewOutgoingContext will
// overwrite any previously-appended metadata.
func NewOutgoing(ctx context.Context, kv map[string]string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.New(kv))
}

// Get returns the value for given key fromt he metadata
func Get(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	key = strings.ToLower(key)

	if kv, ok := md[key]; ok {
		return kv[0]
	}

	return ""
}

// GetOrDefault returns the value for given key fromt he metadata
func GetOrDefault(ctx context.Context, key, fallback string) string {
	// get the value
	value := Get(ctx, key)

	// fallback if the value is empty
	if value == "" {
		value = fallback
	}

	return value
}
