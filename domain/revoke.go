package domain

// deleteAccessToken function removes the specified access
// token from the redis database.
func deleteAccessToken(accessTokenUUID string) error {
	_, err := client.Del(accessTokenUUID).Result()
	return err
}

// deleteRefreshToken function removes the specified refresh
// token from the redis database.
func deleteRefreshToken(refreshTokenUUID string) error {
	_, err := client.Del(refreshTokenUUID).Result()
	return err
}

// Revoke function deletes the specified access & refresh tokens,
// if they exist. It will return nil as error if they didn't exist,
// but if the operation had problem it will return an error.
func Revoke(accessToken, refreshToken string) error {
	// 1. Parse access token.
	_, _, accessTokenUUID, err := parseToken(accessToken)
	if err != nil {
		return err
	}

	// 2. Remove access token from the redis database.
	if err := deleteAccessToken(accessTokenUUID); err != nil {
		return ErrRevokeTokens
	}

	// 3. Parse refresh token.
	_, _, refreshTokenUUID, err := parseToken(refreshToken)
	if err != nil {
		return err
	}

	// 4. Remove refresh token from the redis database.
	if err := deleteRefreshToken(refreshTokenUUID); err != nil {
		return ErrRevokeTokens
	}

	// 5. Return values.
	return nil
}
