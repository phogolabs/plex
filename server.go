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

// OnErrorxHandler can be used to handle an error as general proto message defined by gRPC.
var OnErrorxHandler = http.OnErrorxHandler

// Server represents a server
type Server struct {
	Addr    string
	Gateway *grpc.Gateway
	Proxy   *http.Proxy
}

// NewServer creates a new server
func NewServer() *Server {
	server := &Server{
		Addr:    ":8080",
		Gateway: grpc.NewGateway(),
		Proxy:   http.NewProxy(),
	}

	return server
}

// ListenAndServe listens on the TCP network address srv.Addr and then
// calls Serve to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
func (srv *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}

	return srv.Serve(listener)
}

// Serve accepts incoming connections on the Listener l, creating a
// new service goroutine for each.
func (srv *Server) Serve(listener net.Listener) (err error) {
	var (
		mux   = cmux.New(listener)
		group = errgroup.Group{}
	)

	group.Go(func() error { return srv.Gateway.Serve(mux) })
	group.Go(func() error { return srv.Proxy.Serve(mux) })
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
	if proxy := srv.Proxy; proxy != nil {
		if err := proxy.Shutdown(ctx); err != nil {
			return err
		}
	}

	if gateway := srv.Gateway; gateway != nil {
		if err := gateway.Shutdown(ctx); err != nil {
			return err
		}
	}

	return nil
}
