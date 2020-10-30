package domain

import "os"

func getSecretKey() string {
	secret := os.Getenv("ACCESS_SECRET")
	if secret == "" {
		secret = "SECRET_KEY"
	}
	return secret
}
