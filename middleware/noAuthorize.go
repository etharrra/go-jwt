package middleware

import (
	"net/http"

	"github.com/etharrra/go-jwt/helper"
	"github.com/gin-gonic/gin"
)

func NoAuthorize(c *gin.Context) {
	// Retrieve the token from the "Authorization" cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil || tokenString == "" {
		// No token found, so proceed to the route
		c.Next()
		return
	}

	token, err := helper.TokenParse(tokenString)
	if err == nil && token.Valid {
		// If token is valid, redirect or block access to the login route
		c.JSON(http.StatusForbidden, gin.H{"message": "You are already logged in"})
		c.Abort()
		return
	}

	// If the token is invalid, proceed to the login route
	c.Next()
}
