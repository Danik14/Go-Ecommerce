package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Danik14/go-shop/database"
	"github.com/Danik14/go-shop/handlers"
	"github.com/Danik14/go-shop/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	app := handlers.NewApplication(
		database.ProductData(database.Client, "Products"),
		database.UserData(database.Client, "Users"),
	)

	fmt.Println(app)

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	// router.Use(middleware.Authentication())

	// router.GET("/addtocart", app.AddToCart())
	// router.GET("removeitem", app.RemoveItem())
	// router.GET("/cartcheckout", app.BuyFromCart())
	// router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(": " + port))
}
