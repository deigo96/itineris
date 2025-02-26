package middleware

import (
	"strings"

	"github.com/deigo96/itineris/app/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authorization(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if len(tokenHeader) < 2 {
			c.AbortWithStatusJSON(401, gin.H{
				"code":    4005,
				"message": "Unauthorized",
			})
			return
		}

		token, err := jwt.Parse(tokenHeader[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			return []byte(config.SecretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"code":    4005,
				"message": "Unauthorized",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("nip", claims["nip"])
			c.Set("id", claims["id"])
			c.Set("role", claims["role"])

			c.Next()
		}
	}
}
