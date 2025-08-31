package middelware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// یک کلید مخفی را تعریف کنید. این کلید باید در هر دو سرویس حساب کاربری و گالری عکس یکسان باشد.
var jwtSecretKey = []byte("your_very_secret_key")

// AuthMiddleware تأیید اعتبار توکن JWT را انجام می دهد.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required."})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format is 'Bearer <token>'."})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// توکن را تأیید می کند.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token."})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			// ** اینجا تبدیل به رشته را انجام می‌دهیم **
			if userID, ok := claims["user_id"].(string); ok {
				// اطلاعات کاربر را به context اضافه می کند.
				c.Set("user_id", userID)
				c.Next()
			} else {
				// اگر user_id در Claims موجود نبود یا از نوع رشته نبود
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID in token is invalid."})
				c.Abort()
			}

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims."})
			c.Abort()
			return
		}
	}
}
