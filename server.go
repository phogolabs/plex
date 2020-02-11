package plex

import (
	"context"
	"errors"
	"net"

	"github.com/phogolabs/flaw"
	"github.com/phogolabs/plex/grpc"
	"github.com/phogolabs/plex/http"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
)

const (
	// ErrClosedConn occurs when the connection is closed
	ErrClosedConn = flaw.ErrorConstant("use of closed network connection")
)

// Server represents a server
type Server struct {
	addr    string
	grpcSrv *grpc.Server
	httpSrv *http.Server
}

// NewServer creates a new server
func NewServer(opts ...ServerOption) *Server {
	server := &Server{
		addr:    ":3009",
		grpcSrv: grpc.NewServer(),
		httpSrv: http.NewServer(),
	}

	for _, opt := range opts {
		opt.Apply(server)
	}

	return server
}

// Address returns the address
func (srv *Server) Address() string {
	return srv.addr
}

// Mux returns the http server
func (srv *Server) Mux() *http.Server {
	return srv.httpSrv
}

// Socket returns the grpc server
func (srv *Server) Socket() *grpc.Server {
	return srv.grpcSrv
}

// ListenAndServe listens and server
func (srv *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", srv.addr)
	if err != nil {
		return err
	}

	var (
		mux   = cmux.New(listener)
		group = errgroup.Group{}
	)

	group.Go(func() error { return srv.grpcSrv.Serve(mux) })
	group.Go(func() error { return srv.httpSrv.Serve(mux) })
	group.Go(func() error { return mux.Serve() })

	if err = group.Wait(); err != nil {
		var errx *net.OpError

		if errors.As(err, &errx) {
			if errx.Err.Error() == ErrClosedConn.Error() {
				err = nil
			}
		}

		return err
	}

	return err
}

// Shutdown shutdowns the server
func (srv *Server) Shutdown(ctx context.Context) error {
	if err := srv.httpSrv.Shutdown(ctx); err != nil {
		return err
	}

	if err := srv.grpcSrv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
