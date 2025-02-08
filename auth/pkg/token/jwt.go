package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Id     int64
	Email  string
	Active bool
}

func Create(id int64, active bool, email, secretKey string) (string, time.Time, error) {
	now := time.Now()
	expAt := now.Add(time.Hour * time.Duration(24))

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expAt),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "auth-service",
			Subject:   fmt.Sprintf("%d", id),
		},
		Id:     id,
		Email:  email,
		Active: active,
	}

	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return "", expAt, err
	}

	return tokenStr, expAt, nil
}

func Validate(tokenStr, secretKey string) (int64, string, bool, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, "", false, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return 0, "", false, fmt.Errorf("invalid token")
	}

	v := jwt.NewValidator()
	if err = v.Validate(claims); err != nil {
		return 0, "", false, fmt.Errorf("invalid token")
	}

	return claims.Id, claims.Email, claims.Active, nil
}
