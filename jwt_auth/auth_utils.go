package main

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/savsgio/go-logger/v4"
)

var jwtSignKey = []byte("TestForFasthttpWithJWT")

type userCredential struct {
	Username []byte `json:"username"`
	Password []byte `json:"password"`
	jwt.RegisteredClaims
}

func generateToken(username []byte, password []byte) (string, time.Time) {
	logger.Debugf("Create new token for user %s", username)

	expireAt := time.Now().Add(1 * time.Minute)

	// Embed User information to `token`
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Username: username,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	})

	// token -> string. Only server knows the secret.
	tokenString, err := newToken.SignedString(jwtSignKey)
	if err != nil {
		logger.Error(err)
	}

	return tokenString, expireAt
}

func validateToken(requestToken string) (*jwt.Token, *userCredential, error) {
	logger.Debug("Validating token...")

	user := &userCredential{}
	token, err := jwt.ParseWithClaims(requestToken, user, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})

	return token, user, err
}
