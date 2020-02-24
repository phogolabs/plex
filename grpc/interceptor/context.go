package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation. This technique
// for defining context keys was copied from Go 1.7's new use of context in net/http.
type contextKey struct {
	name string
}

// String represent the context key
func (k *contextKey) String() string {
	return "log: context value " + k.name
}

// ServerStream defines the server-side behavior of a streaming RPC.
//
// All errors returned from ServerStream methods are compatible with the
// status package.
type ServerStream struct {
	Ctx    context.Context
	Stream grpc.ServerStream
}

// SetHeader sets the header metadata. It may be called multiple times.
// When call multiple times, all the provided metadata will be merged.
// All the metadata will be sent out when one of the following happens:
//  - ServerStream.SendHeader() is called;
//  - The first response is sent out;
//  - An RPC status is sent out (error or success).
func (l *ServerStream) SetHeader(m metadata.MD) error {
	return l.Stream.SetHeader(m)
}

// SendHeader sends the header metadata.
// The provided md and headers set by SetHeader() will be sent.
// It fails if called multiple times.
func (l *ServerStream) SendHeader(m metadata.MD) error {
	return l.Stream.SendHeader(m)
}

// SetTrailer sets the trailer metadata which will be sent with the RPC status.
// When called more than once, all the provided metadata will be merged.
func (l *ServerStream) SetTrailer(m metadata.MD) {
	l.Stream.SetTrailer(m)
}

// Context returns the context for this stream.
func (l *ServerStream) Context() context.Context {
	return l.Ctx
}

// SendMsg sends a message. On error, SendMsg aborts the stream and the
// error is returned directly.
//
// SendMsg blocks until:
//   - There is sufficient flow control to schedule m with the transport, or
//   - The stream is done, or
//   - The stream breaks.
//
// SendMsg does not wait until the message is received by the client. An
// untimely stream closure may result in lost messages.
//
// It is safe to have a goroutine calling SendMsg and another goroutine
// calling RecvMsg on the same stream at the same time, but it is not safe
// to call SendMsg on the same stream in different goroutines.
func (l *ServerStream) SendMsg(m interface{}) error {
	return l.Stream.SendMsg(m)
}

// RecvMsg blocks until it receives a message into m or the stream is
// done. It returns io.EOF when the client has performed a CloseSend. On
// any non-EOF error, the stream is aborted and the error contains the
// RPC status.
//
// It is safe to have a goroutine calling SendMsg and another goroutine
// calling RecvMsg on the same stream at the same time, but it is not
// safe to call RecvMsg on the same stream in different goroutines.
func (l *ServerStream) RecvMsg(m interface{}) error {
	return l.Stream.RecvMsg(m)
}
