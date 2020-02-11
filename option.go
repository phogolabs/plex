package plex

// ServerOption represents a server option
type ServerOption interface {
	// Apply applies the option
	Apply(*Server)
}

// ServerOptionFunc represents a server option as function
type ServerOptionFunc func(*Server)

// Apply applies the option
func (op ServerOptionFunc) Apply(srv *Server) {
	op(srv)
}

// WithAddress sets the server address
func WithAddress(addr string) ServerOption {
	fn := func(srv *Server) {
		srv.addr = addr
	}

	return ServerOptionFunc(fn)
}
