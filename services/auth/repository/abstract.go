package repository

import (
	"go-microservices.org/core/utils"
	"go-microservices.org/services/auth/repository/postgres"

	cache "github.com/patrickmn/go-cache"
)

// Repository interface
type Repository interface{}

// AbstractRepository ..
type AbstractRepository struct {
	Repository
	Cache *cache.Cache
}

// NewRepository ...
func NewRepository() *AbstractRepository {
	newRepo := &AbstractRepository{
		Cache: utils.GetCacheHandler(),
	}

	if utils.GetDatasourceInfo() == utils.Postgres {
		newRepo.Repository = &postgres.Repository{
			Connection: utils.GetPostgresHandler(),
		}
	} else {
		newRepo.Repository = nil
	}

	return newRepo
}
