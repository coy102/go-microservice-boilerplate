package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go-microservices.org/api/service"
	servicemodel "go-microservices.org/core/proto"
)

const (
	emptyToken  = "empty token"
	invalidAuth = "invalid authorization type"
)

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ctxKey string

const tokenCtx ctxKey = "token_context"

// AuthCheck ...
func AuthCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// should check client api key first
		// ---

		// skip token check for public
		if isPublicEndpoint(r) {
			next.ServeHTTP(w, r)
			return
		}

		// validate token & session
		authData, err := validateTokenSession(r)
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, buildFailedResponse(err.Error()))
			return
		}

		if !authData.Success {
			writeResponse(w, http.StatusUnauthorized, buildFailedResponse(authData.Message))
			return
		}

		// add token to request so can be used for validation later
		newReq := r.WithContext(addTokenContext(r.Context(), authData))
		next.ServeHTTP(w, newReq)
	})
}

func addTokenContext(ctxParent context.Context, authData *servicemodel.AuthData) context.Context {
	return context.WithValue(ctxParent, tokenCtx, authData)
}

// GetTokenContext ...
func GetTokenContext(ctx context.Context) *servicemodel.AuthData {
	authData, _ := ctx.Value(tokenCtx).(*servicemodel.AuthData)
	return authData
}

func isPublicEndpoint(r *http.Request) bool {
	uri := r.RequestURI
	return (uri == "/login" || uri == "/")
}

func validateTokenSession(r *http.Request) (*servicemodel.AuthData, error) {
	token, err := getTokenFromRequest(r)
	if err != nil {
		return &servicemodel.AuthData{
			Message: err.Error(),
		}, nil
	}

	return service.GetValidateTokenSession(token)
}

func getTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, " ")

	if len(splitToken) == 1 {
		return "", fmt.Errorf(emptyToken)
	}

	if splitToken[0] != "Bearer" {
		return "", fmt.Errorf(invalidAuth)
	}

	return splitToken[1], nil
}

func buildFailedResponse(message string) response {
	return response{
		Message: message,
	}
}

func writeResponse(w http.ResponseWriter, code int, payload response) {
	wresponse, _ := json.Marshal(payload)

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", strings.Join(getAllowedOrigins(), ","))
	w.Header().Add("Access-Control-Allow-Headers", strings.Join(getAllowedHeaders(), ","))
	w.Header().Add("Access-Control-Allow-Credentials", strconv.FormatBool(isAllowCredentials()))
	w.WriteHeader(code)
	w.Write(wresponse)
}
