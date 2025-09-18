package router

import (
	"github.com/gin-gonic/gin"
	auth "github.com/minhanhbb/ecom-golang/app/Controller/Auth"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	v1 := r.Group("/api/v1")

	authGroup := v1.Group("/auth")
	authGroup.POST("/register", auth.Register)
	authGroup.POST("/login", auth.Login)
	authGroup.POST("/logout", auth.Logout)
	authGroup.GET("/profile", auth.JWTAuthMiddleware(), auth.Profile)

	// productGroup := v1.Group("/product")

	// cartGroup := v1.Group("/cart")

	return r
}
