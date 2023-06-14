package middleware

import (
	"os"
	"strings"

	"github.com/rs/cors"
)

// GetCorsAPI ...
func GetCorsAPI() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   getAllowedOrigins(),
		AllowedHeaders:   getAllowedHeaders(),
		AllowCredentials: isAllowCredentials(),
		Debug:            false,
	})
}

func getAllowedOrigins() []string {
	allowedOrigins := []string{"*"}
	if origins := os.Getenv("ALLOWED_ORIGINS"); origins != "" {
		allowedOrigins = strings.Split(origins, "|")
	}

	return allowedOrigins
}

func getAllowedHeaders() []string {
	return []string{
		"Authorization",
		"Content-Type",
	}
}

func isAllowCredentials() bool {
	return true
}
