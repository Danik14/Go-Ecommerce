package routes

import (
	"fmt"

	"github.com/Danik14/go-shop/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	fmt.Println(222222)
	incomingRoutes.POST("/users/signup", handlers.SignUp())
	// incomingRoutes.POST("/users/login", handlers.Login())
	// incomingRoutes.POST("/admin/addproduct", handlers.ProductViewerAdmin())
	// incomingRoutes.GET("/users/productview", handlers.SearchProduct())
	// incomingRoutes.GET("/users/search", handlers.SearchProductByQuery())
}
