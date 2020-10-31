package rest

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// SetRoutes ...
func SetRoutes() {
	router.GET("/refresh", refresh)
	router.GET("/create", create)
	router.GET("/verify", verify)
	router.GET("/revoke", revoke)
	router.Run(":8080")
}
