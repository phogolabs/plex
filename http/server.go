package http

import (
	"context"
	"net/http"
	"reflect"

	"github.com/go-chi/chi"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
)

// Server represents a http server
type Server struct {
	httpSrv *http.Server
	mux     *runtime.ServeMux
}

// NewServer creates a new http server
func NewServer() *Server {
	mux := runtime.NewServeMux()

	router := chi.NewRouter()
	router.Mount("/", mux)

	return &Server{
		mux: mux,
		httpSrv: &http.Server{
			Handler: router,
		},
	}
}

// Router returns the underlying router
func (srv *Server) Router() chi.Router {
	return srv.httpSrv.Handler.(chi.Router)
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
