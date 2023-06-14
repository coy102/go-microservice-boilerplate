package usecase

import (
	"go-microservices.org/core/cache"
	"go-microservices.org/services/auth/repository"
)

// UseCase ...
type UseCase struct {
	repository *repository.AbstractRepository
	session    cache.ICache
}

// NewUseCase ...
func NewUseCase(
	repo *repository.AbstractRepository,
	sess cache.ICache,
) *UseCase {
	return &UseCase{
		repository: repo,
		session:    sess,
	}
}
