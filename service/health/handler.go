package health

import (
	"context"
	"fmt"

	"github.com/phogolabs/log"
	"github.com/phogolabs/plex"
)

var _ HeartbeatAPIServer = &HeartbeatAPI{}

// HeartbeatAPI is the server API for HeartbeatAPI service.
type HeartbeatAPI struct {
	LiveCheckers  []Checker
	ReadyCheckers []Checker
}

// New creates a new HeartbeatAPI
func New() *HeartbeatAPI {
	return &HeartbeatAPI{}
}

// LogLevel returns the desired log level
func (h *HeartbeatAPI) LogLevel() log.Level {
	return log.ErrorLevel
}

// Mount mounts the handler
func (h *HeartbeatAPI) Mount(server *plex.Server) {
	// register the http proxy
	server.Proxy.Use(RegisterHeartbeatAPIHandler)

	// register the server handler
	RegisterHeartbeatAPIServer(server.Gateway.Server, h)
}

// CheckLive checks the live state
func (h *HeartbeatAPI) CheckLive(ctx context.Context, _ *CheckLiveRequest) (*CheckLiveResponse, error) {
	for _, checker := range h.LiveCheckers {
		if err := checker.Check(ctx); err != nil {
			return nil, fmt.Errorf("%s: %w", checker.Name(), err)
		}
	}

	return &CheckLiveResponse{}, nil
}

// CheckReady checks the ready state
func (h *HeartbeatAPI) CheckReady(ctx context.Context, _ *CheckReadyRequest) (*CheckReadyResponse, error) {
	for _, checker := range h.ReadyCheckers {
		if err := checker.Check(ctx); err != nil {
			return nil, err
		}
	}

	return &CheckReadyResponse{}, nil
}

// UseReadyCheck appends a ready check
func (h *HeartbeatAPI) UseReadyCheck(checker Checker) {
	h.ReadyCheckers = append(h.ReadyCheckers, checker)
}

// UseLiveCheck appends a live check
func (h *HeartbeatAPI) UseLiveCheck(checker Checker) {
	h.LiveCheckers = append(h.LiveCheckers, checker)
}
