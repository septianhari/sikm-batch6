package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("session_token")

		if err != nil {
			if ctx.GetHeader("Content-type") == "unauthorized" {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			} else {
				ctx.Redirect(http.StatusSeeOther, "/login")
			}
			return
		}

		// Parsing token JWT
		//tokenString := strings.TrimPrefix(cookie, "Bearer ")
		claims, err := ParseJWT(cookie)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("email", claims.Email)

		ctx.Next()
	})
}

func ParseJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil // Ganti "model.JwtKey" dengan secret key JWT Anda
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// Claims adalah struct untuk menyimpan claim token JWT
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
