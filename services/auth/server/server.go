package main

import (
	"github.com/joho/godotenv"
	"github.com/micro/go-micro/v2/util/log"

	auth "go-microservices.org/core/proto"
	health "go-microservices.org/core/proto/health"
	"go-microservices.org/core/server"
	"go-microservices.org/core/utils"
	"go-microservices.org/services/auth/handler"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found ..")
	}

	log.Info("Data Source : " + utils.GetDatasourceInfo())
}

func main() {
	service := server.NewGRPCServer(server.AuthService)
	service.Init()

	srv := service.Server()
	auth.RegisterAuthHandler(srv, new(handler.Handler))
	health.RegisterHealthHandler(srv, server.NewHealthCheck())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
