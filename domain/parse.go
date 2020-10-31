package domain

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// keyFunc returns a secret key based on the signing method.
func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("error: unexpected signing method")
	}
	return []byte(getSecretKey()), nil
}

// parseToken parses the token and returns the claims.
func parseToken(tokenStr string) (int64, string, string, error) {
	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return 0, "", "", ErrParseToken
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", "", ErrParseToken
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return 0, "", "", ErrParseToken
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return 0, "", "", ErrParseToken
	}
	uuid, ok := claims["uuid"].(string)
	if !ok {
		return 0, "", "", ErrParseToken
	}
	return int64(exp), userID, uuid, nil
}
