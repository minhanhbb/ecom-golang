package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
)

func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("user")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			return
		}
		claimsMap, ok := claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}
		email, ok := claimsMap["email"].(string)
		if !ok || email == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Email not found in token"})
			return
		}
		var user models.User
		if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		if user.IsAdmin != 1 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			return
		}
		c.Next()
	}
}
