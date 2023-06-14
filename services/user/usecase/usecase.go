package usecase

import (
	"go-microservices.org/services/user/client"
	"go-microservices.org/services/user/repository"
)

// UseCase ...
type UseCase struct {
	repository *repository.AbstractRepository
	client     *client.Client
}

// NewUseCase ...
func NewUseCase(repo *repository.AbstractRepository, cl *client.Client) *UseCase {
	return &UseCase{
		repository: repo,
		client:     cl,
	}
}
