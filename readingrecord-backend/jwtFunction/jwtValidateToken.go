package jwtFunction

import (
	"errors"
	"time"

	"github.com/ctampan/ReadingRecordWeb/readingrecord-backend/model"
	"github.com/dgrijalva/jwt-go"
)

func JwtValidateToken(tokenFromHeader string) (string, string, error) {

	token, err := jwt.ParseWithClaims(
		tokenFromHeader,
		&model.JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("somesigningkey"), nil
		},
	)

	if err != nil {
		return "", "", errors.New("couldn't parse token")
	}

	claims, ok := token.Claims.(*model.JwtClaims)
	if !ok {
		return "", "", errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return "", "", errors.New("jwt is expired")
	}

	username := claims.Username

	newToken, err := JwtGenerateToken(username)
	if err != nil {
		return "", "", errors.New("generate new token error")
	}

	return username, newToken, nil
}
