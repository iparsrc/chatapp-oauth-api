package domain

import (
	"errors"
	"time"

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
		return 0, "", "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", "", errors.New("error: can't parse claims")
	}
	exp, ok := claims["exp"].(int64)
	if !ok {
		return 0, "", "", errors.New("error: can't parse exp")
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return 0, "", "", errors.New("error: can't parse uesr_id")
	}
	uuid, ok := claims["uuid"].(string)
	if !ok {
		return 0, "", "", errors.New("error: can't parse uuid")
	}
	return exp, userID, uuid, nil
}

// checkExipirationDate validates that the token is not expired.
func checkExpirationDate(exp int64) error {
	now := time.Now().UTC().Unix()
	if exp < now {
		return errors.New("error: token is expired")
	}
	return nil
}

// checkWhiteList validates that the token is avaliable in redis.
func checkWhiteList(uuid string) error {
	_, err := client.Get(uuid).Result()
	if err != nil {
		return err
	}
	return nil
}

// Verify function returns user_id if the token is valid,
// else it will return an empty string with an error. Also it
// will return empty string & error if an operation goes wrong.
func Verify(tokenStr string) (userID, uuid string, e error) {
	// 1. Parse the token, and get claims.
	exp, userID, uuid, err := parseToken(tokenStr)
	if err != nil {
		return "", "", err
	}

	// 2. Check the expiration date.
	if err := checkExpirationDate(exp); err != nil {
		return "", "", err
	}

	// 3. Check the while list (redis).
	if err := checkWhiteList(uuid); err != nil {
		return "", "", err
	}

	// 4. Return values.
	return userID, uuid, nil
}
