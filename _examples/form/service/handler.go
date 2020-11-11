package service

import (
	"context"
	"fmt"

	"github.com/phogolabs/plex"
	"github.com/phogolabs/plex/_examples/form/sdk"
)

// FooHandler represents a handler
type FooHandler struct{}

// Mount mounts the handler
func (h *FooHandler) Mount(server *plex.Server) {
	// register the http proxy
	server.Proxy.Use(sdk.RegisterFooAPIHandler)

	// register the server handler
	sdk.RegisterFooAPIServer(server.Gateway.Server, h)
}

// Post handles form post requests
func (h *FooHandler) Post(ctx context.Context, input *sdk.FooRequest) (*sdk.FooResponse, error) {
	return &sdk.FooResponse{
		Body: fmt.Sprintf("Welcome, %s!", input.Name),
	}, nil
}
