package handler

import (
	"sync"

	"go-microservices.org/services/user/client"
	"go-microservices.org/services/user/repository"
	"go-microservices.org/services/user/usecase"
)

var uc *usecase.UseCase
var oneUc sync.Once

// GetUsecase ...
func GetUsecase() *usecase.UseCase {
	oneUc.Do(func() {
		uc = usecase.NewUseCase(
			getRepository(),
			getClient(),
		)
	})

	return uc
}

var repo *repository.AbstractRepository
var oneRepo sync.Once

func getRepository() *repository.AbstractRepository {
	oneRepo.Do(func() {
		repo = repository.NewRepository()
	})

	return repo
}

var cl *client.Client
var oneClient sync.Once

func getClient() *client.Client {
	oneClient.Do(func() {
		cl = client.NewClient()
	})

	return cl
}
