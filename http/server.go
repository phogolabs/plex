package http

import (
	"context"
	"net/http"
	"reflect"

	"github.com/go-chi/chi"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
)

// ServeMuxOption is an option that can be given to a ServeMux on construction.
type ServeMuxOption = runtime.ServeMuxOption

// Server represents a http server
type Server struct {
	httpSrv *http.Server
	router  chi.Router
	mux     *runtime.ServeMux
}

// NewServer creates a new http server
func NewServer(opts ...ServeMuxOption) *Server {
	return &Server{
		mux:     runtime.NewServeMux(opts...),
		router:  chi.NewRouter(),
		httpSrv: &http.Server{},
	}
}

// Router returns the underlying router
func (srv *Server) Router() chi.Router {
	return srv.router
}

// Register register a service
func (srv *Server) Register(registrator, service interface{}) {
	fn := reflect.ValueOf(registrator)

	if fn.Type().Kind() != reflect.Func {
		panic("registration must be a function")
	}

	params := []reflect.Value{
		reflect.ValueOf(context.Background()),
		reflect.ValueOf(srv.mux),
		reflect.ValueOf(service),
	}

	fn.Call(params)
}

// Serve serves the mux
func (srv *Server) Serve(mux cmux.CMux) error {
	if srv.httpSrv.Handler == nil {
		srv.router.Mount("/", srv.mux)
		srv.httpSrv.Handler = srv.router
	}

	listener := mux.Match(cmux.HTTP1Fast())
	return srv.httpSrv.Serve(listener)
}

// Shutdown shutdowns the server
func (srv *Server) Shutdown(ctx context.Context) error {
	if err := srv.httpSrv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
