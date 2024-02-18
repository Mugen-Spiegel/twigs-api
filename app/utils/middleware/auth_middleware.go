package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func AuthMiddelware() gin.HandlerFunc {
	return func(request *gin.Context) {
		var authorization_header string = request.GetHeader("Authorization")
		tokenString := strings.Split(authorization_header, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			panic("Invalid Token 1")
		}

		if !token.Valid {
			panic("Invalid Token 2")
		}

		request.Set("current_user_id", "1")
		request.Set("current_user_name", "jophat")
		request.Set("current_user_last_name", "tamayo")
		request.Set("subdivision_id", "tamayo")

		request.Next()
	}
}
