package jwtFunction

import (
	"errors"
	"time"

	"github.com/ctampan/ReadingRecordWeb/readingrecord-backend/model"
	"github.com/dgrijalva/jwt-go"
)

func JwtGenerateToken(username string, password string) (string, error) {
	/*TODO Implement check user credential
	...
	*/

	claims := model.JwtClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Hour * 3).Unix(),
			Issuer:    "someissuer",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("somesigningkey"))

	if err != nil {
		return "", errors.New("jwt is expired")
	}

	return signedToken, nil
}
