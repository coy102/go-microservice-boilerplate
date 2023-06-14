package server

import (
	"context"

	health "go-microservices.org/core/proto/health"
)

// HealthHandler ...
type HealthHandler struct{}

// NewHealthCheck ...
func NewHealthCheck() *HealthHandler {
	return new(HealthHandler)
}

// Check Call is a single request handler called via client.Call or the generated client code
func (*HealthHandler) Check(
	ctx context.Context, in *health.HealthCheckRequest, out *health.HealthCheckResponse,
) error {
	out.Status = health.HealthCheckResponse_SERVING
	return nil
}

// Watch Call is a single request handler called via client.Call or the generated client code
func (*HealthHandler) Watch(
	ctx context.Context, in *health.HealthCheckRequest, stream health.Health_WatchStream,
) error {
	return nil
}
