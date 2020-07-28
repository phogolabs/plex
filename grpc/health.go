package grpc

import (
	"context"

	"github.com/phogolabs/log"
	"github.com/phogolabs/plex/grpc/meta"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

var (
	// StatusNotServing status
	StatusNotServing = &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
	}

	// StatusServing status
	StatusServing = &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}
)

// Checker represents a checker
type Checker interface {
	Check(context.Context) error
}

// Heartbeat represents a health endpoint
type Heartbeat struct {
	ready map[string]Checker
	live  map[string]Checker
}

// NewHeartbeat creates a new heartbeat
func NewHeartbeat() *Heartbeat {
	return &Heartbeat{}
}

// Mount mounts the heartbeat
func (h *Heartbeat) Mount(gw *Gateway) {
	grpc_health_v1.RegisterHealthServer(gw.Server, h)
}

// WithLive adds a liveness check
func (h *Heartbeat) WithLive(name string, check Checker) *Heartbeat {
	if h.live == nil {
		h.live = make(map[string]Checker)
	}

	h.live[name] = check
	// keep handling
	return h
}

// WithReady adds a readiness check
func (h *Heartbeat) WithReady(name string, check Checker) *Heartbeat {
	if h.ready == nil {
		h.ready = make(map[string]Checker)
	}

	h.ready[name] = check
	// keep handling
	return h
}

// Check the service
func (h *Heartbeat) Check(ctx context.Context, r *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	var (
		kv  map[string]Checker
		key = meta.Get(ctx, "X-Gateway-Heartbeat")
	)

	switch key {
	case "ready":
		kv = h.ready
	case "live":
		kv = h.live
	default:
		return nil, status.Errorf(codes.InvalidArgument, "%v not supported", key)
	}

	if kv == nil {
		return StatusServing, nil
	}

	logger := log.GetContext(ctx)

	if name := r.Service; name != "" {
		checker, ok := kv[name]
		if !ok {
			return nil, status.Errorf(codes.NotFound, "%v not found", name)
		}

		kv = make(map[string]Checker)
		kv[name] = checker
	}

	for name, checker := range kv {
		if err := checker.Check(ctx); err != nil {
			logger.WithError(err).Errorf("error checking %v occured", name)
			// return the status
			return StatusNotServing, nil
		}
	}

	return StatusServing, nil
}

// Watch performs a watch for the serving status of the requested service.
func (h *Heartbeat) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "watch method is not supported")
}
