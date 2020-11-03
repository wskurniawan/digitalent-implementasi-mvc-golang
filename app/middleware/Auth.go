package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mitchellh/mapstructure"
)

func Auth(c *gin.Context)  {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return  nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return  []byte("secret"), nil
	})

	if err != nil {
		result := gin.H{
			"message": "invalid token",
			"error": err.Error(),
		}

		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
		return
	}

	fmt.Println("Token verified")
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)

	var idAccount int
	err = mapstructure.Decode(claims["account_number"], &idAccount)
	if err != nil {
		result := gin.H{
			"message": err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}

	fmt.Println(idAccount)
	c.Set("account_number", idAccount)
}
