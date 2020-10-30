package domain

import (
	"errors"
	"time"
)

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
