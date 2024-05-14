package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// Ambil cookie dengan nama "session_token"
		cookie, err := ctx.Cookie("session_token")
		if err != nil {
			if ctx.GetHeader("Content-Type") == "application/json" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			} else {
				ctx.Redirect(http.StatusFound, "/login")
			}
			ctx.Abort()
			return
		}

		// Parsing JWT token pada cookie
		claims := &Claims{}
		jwtKey := []byte("your_secret_key") // Replace with your secret key

		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		// Simpan nilai UserID dari claims ke dalam context dengan key "id"
		ctx.Set("id", claims.UserID)

		// Lanjutkan ke handler berikutnya
		ctx.Next()
	})
}
