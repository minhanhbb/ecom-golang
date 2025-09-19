package router

import (
	"github.com/gin-gonic/gin"
	banners "github.com/minhanhbb/ecom-golang/app/Controller/Admin/Banner"
	category "github.com/minhanhbb/ecom-golang/app/Controller/Admin/Category"
	products "github.com/minhanhbb/ecom-golang/app/Controller/Admin/Product"
	auth "github.com/minhanhbb/ecom-golang/app/Controller/Auth"
	order "github.com/minhanhbb/ecom-golang/app/Controller/Client/Order"
	order_item "github.com/minhanhbb/ecom-golang/app/Controller/Client/Order/OrderItem"
	mw "github.com/minhanhbb/ecom-golang/middleware"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	v1 := r.Group("/api/v1")

	authGroup := v1.Group("/auth")
	authGroup.POST("/register", auth.Register)
	authGroup.POST("/login", auth.Login)
	authGroup.POST("/logout", auth.Logout)
	authGroup.GET("/profile", auth.JWTAuthMiddleware(), mw.AuthRequired(), auth.Profile)

	adminGroup := v1.Group("/admin", auth.JWTAuthMiddleware(), mw.AuthRequired(), mw.AdminRequired())
	categoryGroup := adminGroup.Group("/category")
	categoryGroup.GET("", category.List)
	categoryGroup.GET(":id", category.Detail)
	categoryGroup.POST("", category.Store)
	categoryGroup.PUT(":id", category.Update)
	categoryGroup.DELETE(":id", category.Delete)

	productGroup := adminGroup.Group("/products")
	productGroup.GET("", products.List)
	productGroup.GET(":id", products.Detail)
	productGroup.POST("", products.Store)
	productGroup.PUT(":id", products.Update)
	productGroup.DELETE(":id", products.Delete)

	bannerGroup := adminGroup.Group("/banners")
	bannerGroup.GET("", banners.List)
	bannerGroup.GET(":id", banners.Detail)
	bannerGroup.POST("", banners.Store)
	bannerGroup.PUT(":id", banners.Update)
	bannerGroup.DELETE(":id", banners.Delete)

	orderItemGroup := v1.Group("/order-items")
	orderItemGroup.POST("", order_item.Store)
	orderItemGroup.PUT(":id", order_item.Update)

	orderGroup := v1.Group("/order", auth.JWTAuthMiddleware(), mw.AuthRequired())
	orderGroup.GET("/detail", order.Detail)
	orderGroup.DELETE("/delete", order.Delete)

	return r
}
