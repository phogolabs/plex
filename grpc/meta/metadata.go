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
