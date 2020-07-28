package http

import (
	"encoding/json"
	"net/http"

	"github.com/phogolabs/log"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
)

// HeaderHeartbeat represents a http header passed to the GRPC health check
const HeaderHeartbeat = "X-Gateway-Heartbeat"

// Heartbeat represents a heartbeat
type Heartbeat struct {
	proxy *Proxy
}

// NewHeartbeat creates a new heartbeat
func NewHeartbeat() *Heartbeat {
	return &Heartbeat{}
}

// Mount mounts the heartbeat
func (h *Heartbeat) Mount(proxy *Proxy) {
	// handle the proxy
	h.proxy = proxy
	// mount
	proxy.router.Get("/heartbeat/ready", h.ready)
	proxy.router.Get("/heartbeat/live", h.live)
}

func (h *Heartbeat) ready(w http.ResponseWriter, r *http.Request) {
	// add metadata
	r.Header.Set(HeaderHeartbeat, "ready")
	// do the check
	h.check(w, r)
}

func (h *Heartbeat) live(w http.ResponseWriter, r *http.Request) {
	// add metadata
	r.Header.Set(HeaderHeartbeat, "live")
	// do the check
	h.check(w, r)
}

// Check checks the service
func (h *Heartbeat) check(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		client = grpc_health_v1.NewHealthClient(h.proxy.conn)
		logger = log.GetContext(ctx)
	)

	// add the metadata in regards to
	// https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md#sending-metadata
	ctx = metadata.AppendToOutgoingContext(ctx,
		HeaderHeartbeat, r.Header.Get(HeaderHeartbeat),
	)

	input := &grpc_health_v1.HealthCheckRequest{
		Service: r.URL.Query().Get("service"),
	}

	// output
	output, err := client.Check(ctx, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch output.Status {
	case grpc_health_v1.HealthCheckResponse_UNKNOWN:
		w.WriteHeader(http.StatusNotFound)
	case grpc_health_v1.HealthCheckResponse_SERVING:
		w.WriteHeader(http.StatusOK)
	case grpc_health_v1.HealthCheckResponse_NOT_SERVING:
		w.WriteHeader(http.StatusInternalServerError)
	default:
		w.WriteHeader(http.StatusOK)
	}

	if xerr := json.NewEncoder(w).Encode(output); xerr != nil {
		logger.WithError(xerr).Error("encoding error occurred")
	}
}
