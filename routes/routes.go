package routes

import "github.com/gin-gonic/gin"

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", handlers.SignUp())
	incomingRoutes.POST("/users/login", handlers.Login())
	incomingRoutes.POST("/admin/addproduct", handlers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", handlers.SearchProduct())
	incomingRoutes.GET("/users/search", handlers.SearchProductByQuery())
}
