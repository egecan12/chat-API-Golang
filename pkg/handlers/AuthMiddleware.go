package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization Header'ı al
		authHeader := c.GetHeader("Authorization")
		log.Println("Authorization Header:", authHeader) // Header'ı logla

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			log.Println("Missing or invalid Authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Token'ı Header'dan ayır
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		log.Println("Token String:", tokenString)

		// JWT_SECRET'i al
		secret := os.Getenv("JWT_SECRET")
		log.Println("JWT_SECRET:", secret)

		// Token'ı doğrula
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// İmza metodunu kontrol et
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Unexpected signing method")
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			log.Printf("Token parsing error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Claims'leri al
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			log.Println("Invalid token claims")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// `user_id`'yi claims'den çıkar
		userID, ok := claims["user_id"].(float64)
		if !ok {
			log.Println("Missing user_id in token claims")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// `user_id`'yi bağlama ekle
		log.Println("Authenticated user ID:", int(userID))
		c.Set("user_id", int(userID))

		c.Next()
	}
}
