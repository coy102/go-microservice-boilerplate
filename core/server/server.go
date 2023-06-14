package server

import (
	"os"

	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
)

const (
	// AuthService service name for auth
	AuthService = "go.micro.srv.dmaa.auth"
	// UserService service name for user
	UserService = "go.micro.srv.dmaa.user"
)

// NewGRPCServer ...
func NewGRPCServer(srvName string) service.Service {
	options := getServerOptions(srvName)
	return grpc.NewService(options...)
}

// IsUseKubernetes ...
func IsUseKubernetes() bool {
	return os.Getenv("MICRO_REGISTRY") == "kubernetes"
}

// GetServerOptions ...
func getServerOptions(srvName string) []service.Option {
	options := []service.Option{
		service.Name(srvName),
		service.Version("latest"),
	}

	if IsUseKubernetes() {
		options = append(options, service.Registry(kubernetes.NewRegistry()))
	}

	return options
}
