package jwtFunction

import "errors"

func JwtConfirmCredential(username string, password string) (string, error) {
	// TODO check credential

	token, err := JwtGenerateToken(username)

	if err != nil {
		return "", errors.New("credential error")
	}

	return token, err
}
