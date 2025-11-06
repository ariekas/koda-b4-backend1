package middelware

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func VerifToken() gin.HandlerFunc{
	godotenv.Load()
	
	// token dari env
	var JWTtoken = os.Getenv("APP_SECRET")
	return func(ctx *gin.Context) {
		// ambil header Authorization
		authHeader := ctx.Request.Header.Get("Authorization")

		// ambil isi dari header Authorization terus potong Bearer, ambil token nya aja
		tokenString, _ := strings.CutPrefix(authHeader, "Bearer ")

		// verifikasi token yang dari user sama token dari env
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(JWTtoken), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Next()	
	}
}

