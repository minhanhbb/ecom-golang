package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	models "github.com/minhanhbb/ecom-golang/app/Models"
	"github.com/minhanhbb/ecom-golang/database"
	"github.com/minhanhbb/ecom-golang/utils"
)

func Profile(c *gin.Context) {
	claims, ok := c.Get("user")
	if !ok {
		utils.SonicJSON(c, http.StatusUnauthorized, gin.H{"error": "No user claims found"})
		return
	}
	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		utils.SonicJSON(c, http.StatusUnauthorized, gin.H{"error": "Invalid claims format"})
		return
	}
	email, ok := claimsMap["email"].(string)
	if !ok || email == "" {
		utils.SonicJSON(c, http.StatusUnauthorized, gin.H{"error": "Email not found in token"})
		return
	}
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		utils.SonicJSON(c, http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	utils.SonicJSON(c, http.StatusOK, gin.H{
		"name":  user.Name,
		"email": user.Email,
	})
}
