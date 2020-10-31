package domain

import "errors"

var (
	// ErrCreateToken can occur when creating tokens.
	ErrCreateToken = errors.New("error: can't create the token")
	// ErrSaveToken can occur when saving token to the redis database.
	ErrSaveToken = errors.New("error: can't save the token")
	// ErrParseToken can occur when parsing the token.
	ErrParseToken = errors.New("error: can't parse the token")
	// ErrRefreshTokens can occur when refreshing the tokens.
	ErrRefreshTokens = errors.New("error: can't refresh the tokens")
	// ErrInvalidToken can occur when the token is out dated or is not in the while list.
	ErrInvalidToken = errors.New("error: token is not valid")
	// ErrRevokeTokens can occur when revoking the tokens.
	ErrRevokeTokens = errors.New("error: can't revoke the tokens")
)
