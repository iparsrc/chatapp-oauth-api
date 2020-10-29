package domain

// AccessToken ...
type AccessToken interface {
	Create(userID string) (string, error)
}

type accessToken struct{}
