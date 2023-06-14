module go-microservices.org/services/user

go 1.16

require (
	github.com/joho/godotenv v1.3.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	go-microservices.org/core v0.0.0
)

replace go-microservices.org/core => ../../core

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
