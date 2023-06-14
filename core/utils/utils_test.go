package utils

import (
	"testing"
	"time"

	core "go-microservices.org/core/proto"

	cache "github.com/patrickmn/go-cache"
)

func TestStringToInt(t *testing.T) {
	t.Run("Run Positive", func(t *testing.T) {

		res, err := StringToInt("0")

		if err != nil {
			t.Errorf("Error %s", err)
		} else {
			t.Logf("Success %d", res)
		}
	})

	t.Run("Run Negative", func(t *testing.T) {

		res, err := StringToInt("x")

		if err != nil {
			t.Errorf("Error %s", err)
		} else {
			t.Logf("Success %d", res)
		}
	})
}

func TestStringToFloat(t *testing.T) {
	t.Run("Run Positive", func(t *testing.T) {

		res, err := StringToFloat64("0.0")

		if err != nil {
			t.Errorf("Error %s", err)
		} else {
			t.Logf("Success %f", res)
		}
	})

	t.Run("Run Negative", func(t *testing.T) {

		res, err := StringToFloat64("x")

		if err != nil {
			t.Errorf("Error %s", err)
		} else {
			t.Logf("Success %f", res)
		}
	})
}

func TestUniqueValue(t *testing.T) {
	t.Run("Run Positive", func(t *testing.T) {

		param := []string{"TEST", "TING", "TEST"}
		res := UniqueValue(param)

		t.Logf("Result %s", res)
	})
}

func TestGetCache(t *testing.T) {

	c := cache.New(5*time.Minute, 10*time.Minute)
	var expected interface{}
	expected = "test phrase.."
	key := GetCacheKey("ambrakadole")
	SetCache(key, expected, c)

	t.Run("Run Positive", func(t *testing.T) {
		result := GetCache(key, c)
		t.Logf("Success %s", result)
	})

	t.Run("Run Negative", func(t *testing.T) {
		result := GetCache("key", c)
		t.Logf("Success %s", result)
	})
}

func TestJwtCreateAccessToken(t *testing.T) {
	t.Run("Run Positive", func(t *testing.T) {
		tokenParam := &core.TokenData{
			UserID: 1,
			Permissions: []*core.AccessPermission{
				&core.AccessPermission{Menu: "monitor", Control: []string{"r"}},
			},
		}

		token, err := JwtCreateAccessToken(tokenParam)
		t.Log(token)
		t.Log(err)
	})
}

func TestJwtCreateRefreshToken(t *testing.T) {
	t.Run("Run Positive", func(t *testing.T) {
		token, err := JwtCreateRefreshToken(1)
		t.Log(token)
		t.Log(err)
	})
}

func TestJwtParseAccessToken(t *testing.T) {
	t.Run("Run Positive", func(t *testing.T) {
		tt := `eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJJRCI6IjciLCJwZXJtaXNzaW9ucyI6W3sibWVudSI6Im1vbml0b3IiLCJjb250cm9sIjpbInIiXX0seyJtZW51IjoibW9uaXRvci5kZXRhaWwiLCJjb250cm9sIjpbInIiXX1dLCJtYnNzRGF0YSI6eyJlbWFpbCI6Im1vbjFAbWJzcy5pZCIsImFyZWFDb2RlIjoiU0tBTCJ9fSwiZXhwIjoxNTk0MDkyMTEzLCJpYXQiOjE1OTQwODg1MTMsImlzcyI6InplYnJheC5pZCJ9.qjVYUC_UevQrhyXMC2TMHfHHytDD8wUTu6XWBvvzK7nw3wH8AyYWWOpiokcKhSoycegx_nPRunN4zSuhsNbibQ`

		auth := JwtParseAccessToken(tt)
		t.Log(auth)
	})
}

func TestJwtParseRefreshToken(t *testing.T) {
	t.Run("Run Positive", func(t *testing.T) {
		tt := `eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJJRCI6IjciLCJwZXJtaXNzaW9ucyI6W3sibWVudSI6Im1vbml0b3IiLCJjb250cm9sIjpbInIiXX0seyJtZW51IjoibW9uaXRvci5kZXRhaWwiLCJjb250cm9sIjpbInIiXX1dLCJtYnNzRGF0YSI6eyJlbWFpbCI6Im1vbjFAbWJzcy5pZCIsImFyZWFDb2RlIjoiU0tBTCJ9fSwiZXhwIjoxNTk0MDg5MTc2LCJpYXQiOjE1OTQwODU1NzYsImlzcyI6InplYnJheC5pZCJ9.qAt4C02avli7NaBRh9j2E9_e-tGctffZqnjB8-qjk6hlvCdwz1m6HLCtCSf3GoCHz-mnobmIQir0tC_SLL9nyA`

		auth := JwtParseRefreshToken(tt)
		t.Log(auth)
	})
}

func TestGetHashPassword(t *testing.T) {
	t.Run("Run Positive", func(t *testing.T) {
		p := `insight2020`

		t.Log(GetHashPassword(p))
	})
}
