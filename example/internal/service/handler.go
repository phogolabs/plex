package service

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/phogolabs/plex"
	"github.com/phogolabs/plex/example/sdk"
)

var _ sdk.UserAPIServer = &UserAPI{}

// UserAPI represents a handler
type UserAPI struct{}

// Mount mounts the handler
func (h *UserAPI) Mount(server *plex.Server) {
	// register the http proxy
	server.Proxy.Use(sdk.RegisterUserAPIHandler)

	// register the server handler
	sdk.RegisterUserAPIServer(server.Gateway.Server, h)
}

// Create a user for given email and password
func (h *UserAPI) CreateUser(ctx context.Context, input *sdk.CreateUserRequest) (*sdk.CreateUserResponse, error) {
	spew.Dump(input)

	response := &sdk.CreateUserResponse{
		Id: "457a1114-8832-4ad2-b950-33eee5fab920",
	}

	return response, nil
}
