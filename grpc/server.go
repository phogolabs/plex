package grpc

import (
	"context"
	"reflect"

	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

const (
	// ContentType is a grpc content-type
	ContentType = "application/grpc"
)

// Interceptor represents a grpc interface
type Interceptor interface {
	Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	Stream(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
}

// Server represents a grpc server
type Server struct {
	grpcSrv *grpc.Server
	chain   ChainInterceptor
}

// NewServer creates a new grpc server
func NewServer() *Server {
	return &Server{}
}

// Use uses an interceptor
func (srv *Server) Use(interceptors ...Interceptor) {
	srv.chain = append(srv.chain, interceptors...)
}

// Register register a service
func (srv *Server) Register(registrator, service interface{}) {
	srv.initialize()
	fn := reflect.ValueOf(registrator)

	if fn.Type().Kind() != reflect.Func {
		panic("registration must be a function")
	}

	params := []reflect.Value{
		reflect.ValueOf(srv.grpcSrv),
		reflect.ValueOf(service),
	}

	fn.Call(params)
}

// Serve serves the mux
func (srv *Server) Serve(mux cmux.CMux) error {
	srv.initialize()
	listener := mux.Match(cmux.HTTP2HeaderField("content-type", ContentType))
	return srv.grpcSrv.Serve(listener)
}

// Shutdown shutdowns the server
func (srv *Server) Shutdown(ctx context.Context) error {
	if srv.grpcSrv == nil {
		return nil
	}

	srv.grpcSrv.GracefulStop()
	return nil
}

func (srv *Server) initialize() {
	if srv.grpcSrv != nil {
		return
	}

	srv.grpcSrv = grpc.NewServer(
		grpc.UnaryInterceptor(srv.chain.Unary),
		grpc.StreamInterceptor(srv.chain.Stream),
	)
}
