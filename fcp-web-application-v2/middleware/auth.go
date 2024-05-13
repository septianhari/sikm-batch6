package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// Get the Authorization header
		authHeader := ctx.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Extract the token from the Authorization header
		token := strings.Split(authHeader, "Bearer ")[1]

		// Verify the token
		jwtKey := []byte(os.Getenv("JWT_KEY"))
		tokenClaims, err := jwt.ParseWithClaims(token, jwt.MapClaims{
			"iss": "fcp-web-application-v2",
			"aud": "fcp-web-application-v2",
		}, jwtKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Check if the token is valid
		if claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && claims["iss"] == "fcp-web-application-v2" && claims["aud"] == "fcp-web-application-v2" {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
	})
}
