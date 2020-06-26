package service

import (
	"github.com/phogolabs/plex"
	"github.com/phogolabs/plex/http/responder"
)

// New creates a new server
func New() *plex.Server {
	// create the plex server
	server := plex.NewServer()
	server.Proxy.OnError(responder.PostgreSQLErrorFormatter)

	// handler setup
	handler := &FooHandler{}
	handler.Mount(server)

	// return the server
	return server
}
