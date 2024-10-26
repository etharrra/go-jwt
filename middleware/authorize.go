package middleware

import (
	"net/http"
	"time"

	"github.com/etharrra/go-jwt/helper"
	"github.com/etharrra/go-jwt/initializers"
	"github.com/etharrra/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authorize(c *gin.Context) {
	// Define a constant for the cookie name
	const AuthCookieName = "Authorization"

	// get the cookie from req using the constant
	tokenString, err := c.Cookie(AuthCookieName)
	if err != nil {
		// If the cookie is missing or empty, respond with Unauthorized
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		c.Abort()
		return
	}

	// decode/validate it
	token, err := helper.TokenParse(tokenString)

	// Log errors for better traceability
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Ensure claims are in the correct format
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token claiming error"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check the expiration time of the token
	exp, ok := claims["exp"].(float64)
	if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Find the user with the subject (sub) claim
	userID, ok := claims["sub"].(float64)
	// fmt.Println(reflect.TypeOf(claims["sub"]))
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "can't claim user id"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "id = ?", userID) // Use "id = ?" to prevent injection
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": result.Error.Error()})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Attach user to the request context
	c.Set("user", user)

	// Continue with the request
	c.Next()
}
