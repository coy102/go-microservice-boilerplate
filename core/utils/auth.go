package utils

import (
	"os"
	"time"

	core "go-microservices.org/core/proto"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaimType int

const (
	invalidIssuer = "invalid issuer"
	invalidToken  = "invalid token data"
)

// AccessTokenClaims ...
type AccessTokenClaims struct {
	Data *core.TokenData `json:"data"`
	*jwt.StandardClaims
}

// RefreshTokenClaims ...
type RefreshTokenClaims struct {
	UserID int32 `json:"userID"`
	*jwt.StandardClaims
}

// JwtCreateAccessToken ...
func JwtCreateAccessToken(tokenParam *core.TokenData) (string, error) {
	expires := GetAccessTokenExpires()
	claims := &AccessTokenClaims{
		Data:           tokenParam,
		StandardClaims: createStandardClaim(expires),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString(getJwtSigningKey())
}

// JwtCreateRefreshToken ...
func JwtCreateRefreshToken(userID int32) (string, error) {
	expires := GetRefreshTokenExpires()
	claims := &RefreshTokenClaims{
		UserID:         userID,
		StandardClaims: createStandardClaim(expires),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString(getJwtSigningKey())
}

func createStandardClaim(expires time.Duration) *jwt.StandardClaims {
	return &jwt.StandardClaims{
		Issuer:    getJwtIssuer(),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: GetTokenExpiresTime(expires),
	}
}

// GetAccessTokenExpires ...
func GetAccessTokenExpires() time.Duration {
	expire, _ := StringToDuration(os.Getenv("TOKEN_EXPIRE_HOUR"))
	return expire
}

// GetRefreshTokenExpires ...
func GetRefreshTokenExpires() time.Duration {
	expire, _ := StringToDuration(os.Getenv("REFRESH_TOKEN_EXPIRE_HOUR"))
	return expire
}

// GetTokenExpiresTime ...
func GetTokenExpiresTime(expires time.Duration) int64 {
	return time.Now().Add(expires * time.Hour).Unix()
}

// JwtParseAccessToken ...
func JwtParseAccessToken(accessToken string) *core.AuthData {
	authData := new(core.AuthData)
	jwtToken, err := jwt.ParseWithClaims(
		accessToken,
		&AccessTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return getJwtSigningKey(), nil
		},
	)

	if err != nil {
		authData.Message = err.Error()
		return authData
	}

	tokenClaim := jwtToken.Claims.(*AccessTokenClaims)
	if err := tokenClaim.Valid(); err != nil {
		authData.Message = err.Error()
		return authData
	}

	if tokenClaim.Data == nil {
		authData.Message = invalidToken
		return authData
	}

	if valid := tokenClaim.VerifyIssuer(getJwtIssuer(), true); !valid {
		authData.Message = invalidIssuer
		return authData
	}

	authData.Success = true
	authData.AccessToken = accessToken
	authData.ParsedToken = tokenClaim.Data
	return authData
}

// JwtParseRefreshToken ...
func JwtParseRefreshToken(refreshToken string) *core.AuthData {
	authData := new(core.AuthData)
	jwtToken, err := jwt.ParseWithClaims(
		refreshToken,
		&RefreshTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return getJwtSigningKey(), nil
		},
	)

	if err != nil {
		authData.Message = err.Error()
		return authData
	}

	tokenClaim := jwtToken.Claims.(*RefreshTokenClaims)
	if err := tokenClaim.Valid(); err != nil {
		authData.Message = err.Error()
		return authData
	}

	if tokenClaim.UserID == 0 {
		authData.Message = invalidToken
		return authData
	}

	if valid := tokenClaim.VerifyIssuer(getJwtIssuer(), true); !valid {
		authData.Message = invalidIssuer
		return authData
	}

	authData.Success = true
	authData.RefreshToken = refreshToken
	authData.ParsedToken = &core.TokenData{UserID: tokenClaim.UserID}
	return authData
}

func getJwtSigningKey() []byte {
	return []byte(os.Getenv("JWT_SIGNING_KEY"))
}

func getJwtIssuer() string {
	return os.Getenv("JWT_ISSUER")
}
