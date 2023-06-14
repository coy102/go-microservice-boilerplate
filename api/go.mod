module go-microservices.org/api

go 1.16

require (
	github.com/99designs/gqlgen v0.17.9
	github.com/go-chi/chi v3.3.2+incompatible
	github.com/gorilla/websocket v1.5.0
	github.com/joho/godotenv v1.3.0
	github.com/rs/cors v1.7.0
	github.com/vektah/gqlparser/v2 v2.4.4
	go-microservices.org/core v0.0.0
)

replace go-microservices.org/core => ../core

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
