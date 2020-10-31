package domain

// deletePreviousRefreshToken removes the token from redis.
func deletePreviousRefreshToken(uuid string) error {
	_, err := client.Del(uuid).Result()
	return err
}

// Refresh function takes a refresh token and creates new tokens.
// It will return empty strings as tokens and an error if the
// refresh token is invalid or something goes wrong.
func Refresh(refreshToken string) (newAccessToken, newRefreshToken string, e error) {
	// 1. Verify the refresh token.
	userID, uuid, err := Verify(refreshToken)
	if err != nil {
		return "", "", err
	}

	// 2. Delete the previous refresh token from the redis.
	if err := deletePreviousRefreshToken(uuid); err != nil {
		return "", "", ErrRefreshTokens
	}

	// 3. Create new tokens.
	newAccessToken, newRefreshToken, err = Create(userID)
	if err != nil {
		return "", "", err
	}

	// 4. Return values.
	return newAccessToken, newRefreshToken, nil
}
