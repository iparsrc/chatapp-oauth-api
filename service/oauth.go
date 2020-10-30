package service

import (
	"github.com/parsaakbari1209/Chatapp-oauth-api/domain"
)

// OAuth interface defines the available service methods.
type OAuth interface {
	Refresh(refreshToken string) (newAccessToken, newRefreshToken string, e error)
	Create(userID string) (accessToken, refreshToken string, e error)
	Verify(tokenStr string) (userID, uuid string, e error)
	Revoke(accessToken, refreshToken string) error
}

// NewOAuth returns an implemented struct of OAuth interface.
func NewOAuth() OAuth {
	return &oauth{}
}

type oauth struct{}

func (oa *oauth) Create(userID string) (accessToken, refreshToken string, e error) {
	accessToken, refreshToken, e = domain.Create(userID)
	return accessToken, refreshToken, e
}

func (oa *oauth) Verify(tokenStr string) (userID, uuid string, e error) {
	userID, uuid, e = domain.Verify(tokenStr)
	return userID, uuid, e
}

func (oa *oauth) Refresh(refreshToken string) (newAccessToken, newRefreshToken string, e error) {
	newAccessToken, newRefreshToken, e = domain.Refresh(refreshToken)
	return newAccessToken, newRefreshToken, e
}

func (oa *oauth) Revoke(accessToken, refreshToken string) error {
	return domain.Revoke(accessToken, refreshToken)
}
