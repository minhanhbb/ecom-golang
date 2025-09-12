package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	// In a real application, you would invalidate a token here.
	// For now, we'll just send a success message.
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
