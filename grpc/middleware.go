package grpc

import (
	"github.com/soheilhy/cmux"
)

const (
	// ContentType is a grpc content-type
	ContentType = "application/grpc"
)

// Match matches the request
var Match = cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", ContentType)
