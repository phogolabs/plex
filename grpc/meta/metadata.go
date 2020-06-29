package meta

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

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
