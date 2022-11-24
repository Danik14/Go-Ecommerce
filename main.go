package main

import (
	"log"
	"os"

	"github.com/Danik14/go-shop/database"
	"github.com/Danik14/go-shop/handlers"
	"github.com/Danik14/go-shop/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//getting a port from env variable
	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	//creating a new application for easer access to collections
	app := handlers.NewApplication(
		database.ProductData(database.Client, "Products"),
		database.UserData(database.Client, "Users"),
	)

	//creating new serve mux
	router := gin.New()

	//using middleware for logging every request
	router.Use(gin.Logger())

	//registering routes related to user
	routes.UserRoutes(router)
	// router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	// router.GET("removeitem", app.RemoveItem())
	// router.GET("/cartcheckout", app.BuyFromCart())
	// router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(": " + port))
}
