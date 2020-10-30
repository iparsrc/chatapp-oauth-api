package domain

// TokenDetails is to move tokens and their details easily.
type TokenDetails struct {
	AccessToken      string
	RefreshToken     string
	AccessTokenUUID  string
	RefreshTokenUUID string
	AccessTokenExp   int64
	RefreshTokenExp  int64
}
