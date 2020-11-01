package rest

import (
	"net/http"

	"github.com/parsaakbari1209/Chatapp-oauth-api/domain"

	"github.com/gin-gonic/gin"
	"github.com/parsaakbari1209/Chatapp-oauth-api/service"
	"github.com/parsaakbari1209/Chatapp-oauth-api/utils"
)

var (
	s = service.NewOAuth()
)

func create(c *gin.Context) {
	// 1. Get user_id from the request url.
	userID := c.Query("user_id")
	if userID == "" {
		restErr := utils.BadRequest("user_id is not specified.")
		c.JSON(restErr.Status, restErr)
		return
	}

	// 2. Create access and refresh tokens.
	accessToken, refreshToken, err := s.Create(userID)
	if err != nil {
		restErr := utils.InternalServerErr("can't operate creation.")
		c.JSON(restErr.Status, restErr)
		return
	}

	// 3. Set access and refresh tokens as cookies.
	c.SetCookie("access_token", accessToken, 900, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, 604800, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{})
	return
}

func refresh(c *gin.Context) {
	// 1. Get refresh token from the cookies.
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		restErr := utils.BadRequest("refresh_token cookie is not set.")
		c.JSON(restErr.Status, restErr)
		return
	}

	// 2. Create new access and refresh tokens.
	newAccessToken, newRefreshToken, err := s.Refresh(refreshToken)
	if err != nil {
		restErr := utils.BadRequest("can't operate refreshing.")
		c.JSON(restErr.Status, restErr)
		return
	}

	// 3. Override the cookies.
	c.SetCookie("access_token", newAccessToken, 900, "/", "localhost", false, true)
	c.SetCookie("refresh_token", newRefreshToken, 604800, "/", "localhost", false, true)

	// 4. Return the results.
	c.JSON(http.StatusOK, gin.H{})
	return
}

func verify(c *gin.Context) {
	// 1. Get access token from cookies.
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		restErr := utils.BadRequest("access_token cookie is not set.")
		c.JSON(restErr.Status, restErr)
		return
	}

	// 2. Verify the access token.
	_, _, err = s.Verify(accessToken)
	if err == domain.ErrParseToken {
		restErr := utils.BadRequest("can't operate verification.")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err == domain.ErrInvalidToken {
		c.JSON(http.StatusOK, gin.H{"is_valid": false})
		return
	}

	// 3. Return the results.
	c.JSON(http.StatusOK, gin.H{"is_valid": true})
	return
}

func revoke(c *gin.Context) {
	// 1. Get the refresh token from cookies.
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		restErr := utils.BadRequest("refresh_token cookie is not set.")
		c.JSON(restErr.Status, restErr)
		return
	}

	// 2. Get the access token from cookies.
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		restErr := utils.BadRequest("access_token cookie is not set.")
		c.JSON(restErr.Status, restErr)
		return
	}

	// 3. Revoke the tokens.
	if err := s.Revoke(accessToken, refreshToken); err != nil {
		restErr := utils.BadRequest("can't operate revokation.")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"succeed": true})
	return
}
