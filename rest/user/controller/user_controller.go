package controller

import (
	"errors"
	"net/http"
	"time"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateToken(secretKey string, data map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(time.Hour * 72).Unix(), // Token expiration time
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (string, error) {
	// Initialize a new instance of `Claims`
	claims := jwt.MapClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well.  This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", err
	}
	data, ok := claims["data"].(map[string]interface{})
	if (!ok) {
		return "", errors.New("Could not extract data")
	}

	user, ok := data["username"].(string)
	if (!ok) {
		return "", errors.New("Could not extract username")
	}
	if (user != "admin") {
		return "", errors.New("Invalid user")
	}

	return user, nil
}

const SecretKey = "abkckefkghiklsdkjflwel2l1l49s"

func Login(c *gin.Context) {
	var userDTO UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userDTO.Username != "admin" || userDTO.Password != "admin1" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	data := map[string]interface{}{
		"username": "admin",
	}
	token, err := CreateToken(SecretKey, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func ProtectedPath(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}
	token := authHeader[7:]
	_, err := ValidateToken(token)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// get user from context
	user, _ := c.Get("user")
	c.String(http.StatusOK, user.(string))
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}
		token := authHeader[7:]
		user , err := ValidateToken(token)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Set("user", user)
		c.Next()
	}
}