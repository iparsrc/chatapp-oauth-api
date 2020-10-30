package domain

import (
	"errors"
	"os"
	"time"

	"github.com/twinj/uuid"

	"github.com/dgrijalva/jwt-go"
)

// TokenDetails ...
type TokenDetails struct {
	AccessToken      string
	RefreshToken     string
	AccessTokenUUID  string
	RefreshTokenUUID string
	AccessTokenExp   int64
	RefreshTokenExp  int64
}

func getSecretKey() string {
	secret := os.Getenv("ACCESS_SECRET")
	if secret == "" {
		secret = "SECRET_KEY"
	}
	return secret
}

// createAccessToken function creates an access token with the
// given credentials and sets it to the given TokenDetails.
// It returns an error if the operation goes wrong.
func createAccessToken(userID string, td *TokenDetails) error {
	atClaims := jwt.MapClaims{
		"uuid":    td.AccessTokenUUID,
		"exp":     td.AccessTokenExp,
		"user_id": userID,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	atStr, err := at.SignedString([]byte(getSecretKey()))
	if err != nil {
		return errors.New("error: can't create access token")
	}
	td.AccessToken = atStr
	return nil
}

// createRefreshToken function creates a refresh token with the
// given credentials and sets it to the given TokenDetails.
// It returns an error if the operations goes wrong.
func createRefreshToken(userID string, td *TokenDetails) error {
	rtClaims := jwt.MapClaims{
		"uuid":    td.RefreshTokenUUID,
		"exp":     td.RefreshTokenExp,
		"user_id": userID,
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rtStr, err := rt.SignedString([]byte(getSecretKey()))
	if err != nil {
		return errors.New("error: can't create refresh token")
	}
	td.RefreshToken = rtStr
	return nil
}

// saveAccessAndRefreshTokens function saves token_uuid/user_id key/value
// pairs inside a redis database, and returns error if it can't save the pairs.
func saveAccessAndRefreshTokens(userID string, td *TokenDetails) error {
	accessTokenExp := time.Unix(td.AccessTokenExp, 0).UTC()
	refreshTokenExp := time.Unix(td.RefreshTokenExp, 0).UTC()
	now := time.Now().UTC()
	if err := client.Set(td.AccessTokenUUID, userID, accessTokenExp.Sub(now)).Err(); err != nil {
		return err
	}
	if err := client.Set(td.RefreshTokenUUID, userID, refreshTokenExp.Sub(now)).Err(); err != nil {
		return err
	}
	return nil
}

// Create function creates access and refresh tokens and,
// saves them in a redis database.
func Create(userID string) (accessToken, refreshToken string, e error) {
	// 1. Create the token details.
	td := &TokenDetails{
		AccessTokenUUID:  uuid.NewV4().String(),
		RefreshTokenUUID: uuid.NewV4().String(),
		AccessTokenExp:   time.Now().Add(time.Minute * 15).Unix(),
		RefreshTokenExp:  time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	// 2. Create access token.
	if err := createAccessToken(userID, td); err != nil {
		return "", "", err
	}

	// 3. Create refresh token.
	if err := createRefreshToken(userID, td); err != nil {
		return "", "", err
	}

	// 4. Save access and refresh tokens in the redis database.
	if err := saveAccessAndRefreshTokens(userID, td); err != nil {
		return "", "", err
	}

	// 5. Return values.
	return td.AccessToken, td.RefreshToken, nil
}
