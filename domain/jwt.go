package domain

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/parsaakbari1209/Chatapp-oauth-api/utils"
)

// AccessToken ...
type AccessToken interface {
	Create(userID string) (string, error)
}

func getSecretKey() string {
	secret := os.Getenv("ACCESS_SECRET")
	if secret == "" {
		secret = "SECRET_KEY"
	}
	return secret
}

type accessToken struct{}

func (t *accessToken) Create(userID string) (string, *utils.RestErr) {
	// 1. Create  token claims map.
	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        time.Now().UTC().Add(time.Minute * 15).Unix(),
	}

	// 2. Prepare token creation.
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	// 3. Create and get token in string.
	token, err := at.SignedString([]byte(getSecretKey()))
	if err != nil {
		return "", utils.InternalServerErr("error: can't create access token")
	}

	// 4. Return values.
	return token, nil
}
