package postgres

import (
	"go-microservices.org/core/connection"
)

// Repository ...
type Repository struct {
	Connection connection.Connection
}
