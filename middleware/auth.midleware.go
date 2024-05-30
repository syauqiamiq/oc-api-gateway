package middleware

import (
	"net/http"
	"ocApiGateway/dto"
	"ocApiGateway/helper"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenRaw := c.GetHeader("Authorization")
		if tokenRaw == "" {
			response := helper.APIResponse(http.StatusUnauthorized, "Authorization header is missing", nil)
			c.JSON(response.Code, response)
			c.Abort()
			return
		}

		splittedToken := strings.Split(tokenRaw, " ")
		if len(splittedToken) != 2 {
			response := helper.APIResponse(http.StatusUnauthorized, "Authorization header format is invalid", nil)
			c.JSON(response.Code, response)
			c.Abort()
			return
		}

		tokenString := splittedToken[1]

		token, err := jwt.ParseWithClaims(tokenString, &dto.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_TOKEN_SECRET")), nil
		})
		if err != nil {
			response := helper.APIResponse(http.StatusUnauthorized, "Invalid token", nil)
			c.JSON(response.Code, response)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*dto.MyClaims); ok && token.Valid {
			// Check if the token is expired
			if claims.ExpiresAt.Time.Before(time.Now()) {
				response := helper.APIResponse(http.StatusUnauthorized, "Token is expired", nil)
				c.JSON(response.Code, response)
				c.Abort()
				return
			}

			// Set the claims in the session
			sess := sessions.Default(c)
			sess.Set("TOKEN-ACCESS-DATA", claims)
			c.Next()
		} else {
			response := helper.APIResponse(http.StatusInternalServerError, "Error parsing JWT", nil)
			c.JSON(response.Code, response)
			c.Abort()
		}
	}
}
func GetSessionAccessData(c *gin.Context) *dto.MyClaims {
	sess := sessions.Default(c)
	return sess.Get("TOKEN-ACCESS-DATA").(*dto.MyClaims)
}
