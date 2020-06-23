package service

import "github.com/phogolabs/plex"

// New creates a new server
func New() *plex.Server {
	// create the plex server
	server := plex.NewServer()

	// handler setup
	handler := &FooHandler{}
	handler.Mount(server)

	// return the server
	return server
}
