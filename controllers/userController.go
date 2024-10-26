package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/etharrra/go-jwt/initializers"
	"github.com/etharrra/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

/**
 * Singup handles the signup process for a new user.
 * It retrieves the email and password from the request body, hashes the password,
 * creates a new user in the database, and returns appropriate JSON responses based on the outcome.
 */
func Singup(c *gin.Context) {
	// get the email & psw from req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// hash the psw
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// create user
	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	// response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}

func Login(c *gin.Context) {
	// get email and psw from req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// look up requested user
	var user models.User
	result := initializers.DB.First(&user, "email = ?", body.Email)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	// check psw
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid E-mail or Password!",
		})

		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRECT")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create token string",
		})

		return
	}

	// set cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// send it back
	c.JSON(http.StatusOK, gin.H{
		"message": "login Success!",
	})
}

func Home(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func Singout(c *gin.Context) {
	fmt.Print("singout")
	// Clear the Authorization cookie by setting it with an expired date
	c.SetCookie("Authorization", "", -1, "/", "", false, true)

	c.Set("user", nil)

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully signed out",
	})
}
