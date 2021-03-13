package runtime

import "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
var HTTPStatusFromCode = runtime.HTTPStatusFromCode
