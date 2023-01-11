package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/phogolabs/restify/middleware"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

// ServeMuxOption is an option that can be given to a ServeMux on construction.
type ServeMuxOption = runtime.ServeMuxOption

// ProxyHandlerFunc handles the proxy call
type ProxyHandlerFunc = func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

// Proxy represents a http proxy to the server
type Proxy struct {
	mux      *runtime.ServeMux
	conn     *grpc.ClientConn
	server   *http.Server
	handlers []ProxyHandlerFunc
	opts     []grpc.DialOption
	onError  runtime.ErrorHandlerFunc
	router   chi.Router
}

// NewProxy creates a new http proxy
func NewProxy(opts ...ServeMuxOption) *Proxy {
	// router setup
	router := chi.NewRouter()

	// middleware by recommendation
	router.Use(middleware.NoCache)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	// middleware extensions
	router.Use(Accept)
	router.Use(ContentType)
	router.Use(Metadata)

	proxy := &Proxy{
		router:  router,
		onError: runtime.DefaultHTTPErrorHandler,
		server: &http.Server{
			Handler: router,
		},
	}

	// mux setup
	opts = append(opts, AllIncomingHeaders)
	opts = append(opts, AllOutgoingHeaders)

	// encoding
	opts = append(opts, WithJSONMarshaler)
	opts = append(opts, WithFormMarshaler)

	// responding
	opts = append(opts, WithForwardResponse)

	// events
	opts = append(opts, proxy.WithErrorHandler())

	// setup the mux
	proxy.mux = runtime.NewServeMux(opts...)

	return proxy
}

// Use registers the proxy handler
func (proxy *Proxy) Use(fn ProxyHandlerFunc) {
	proxy.handlers = append(proxy.handlers, fn)
}

// UseDialOption uses specific dial option
func (proxy *Proxy) UseDialOption(opt grpc.DialOption) {
	proxy.opts = append(proxy.opts, opt)
}

// OnError handles response errors
func (proxy *Proxy) OnError(fn runtime.ErrorHandlerFunc) {
	proxy.onError = fn
}

// Serve serves the mux
func (proxy *Proxy) Serve(mux cmux.CMux) error {
	listener := mux.Match(Match)

	if err := proxy.connect(listener.Addr()); err != nil {
		return err
	}

	if err := proxy.attach(); err != nil {
		return err
	}

	return proxy.server.Serve(listener)
}

// Shutdown shutdowns the server
func (proxy *Proxy) Shutdown(ctx context.Context) error {
	if err := proxy.server.Shutdown(ctx); err != nil {
		return err
	}

	if conn := proxy.conn; conn != nil {
		return conn.Close()
	}

	return nil
}

// Router returns the underlying router
func (proxy *Proxy) Router() chi.Router {
	return proxy.router
}

func (proxy *Proxy) connect(addr net.Addr) error {
	address, err := proxy.address(addr)
	if err != nil {
		return err
	}

	params := grpc.ConnectParams{
		MinConnectTimeout: 5 * time.Minute,
		Backoff:           backoff.DefaultConfig,
	}

	opts := proxy.opts
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithReturnConnectionError())
	opts = append(opts, grpc.WithConnectParams(params))

	if proxy.conn, err = grpc.Dial(address, opts...); err != nil {
		return err
	}

	return nil
}

func (proxy *Proxy) address(addr net.Addr) (string, error) {
	_, port, err := net.SplitHostPort(addr.String())
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("127.0.0.1:%s", port), nil
}

func (proxy *Proxy) attach() error {
	proxy.router.Mount("/", proxy.mux)

	for _, fn := range proxy.handlers {
		if err := fn(context.Background(), proxy.mux, proxy.conn); err != nil {
			return err
		}
	}

	return nil
}

// WithErrorHandler creates a error handler proxy
func (proxy *Proxy) WithErrorHandler() runtime.ServeMuxOption {
	fn := func(ctx context.Context, mux *runtime.ServeMux, m runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
		if proxy.onError != nil {
			proxy.onError(ctx, mux, m, w, r, err)
		}
	}

	return runtime.WithErrorHandler(fn)
}
