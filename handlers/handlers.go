package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Danik14/go-shop/database"
	"github.com/Danik14/go-shop/models"
	"github.com/Danik14/go-shop/tokens"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var Validate = validator.New()

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		fmt.Println(1)

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "This phone is alredy used"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		//RFC3339 is a specified time format
		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refresh_token, _ := tokens.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refresh_token
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		_, insertErr := UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Some server error"})
			return
		}
		defer cancel()

		c.JSON(http.StatusCreated, "Successfuly Created")
	}
}

// func Login() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		defer cancel()
// 		var user models.User
// 		var founduser models.User
// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 			return
// 		}
// 		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
// 			return
// 		}
// 		PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
// 		defer cancel()
// 		if !PasswordIsValid {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 			fmt.Println(msg)
// 			return
// 		}
// 		token, refreshToken, _ := tokens.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, founduser.User_ID)
// 		defer cancel()
// 		tokens.UpdateAllTokens(token, refreshToken, founduser.User_ID)
// 		c.JSON(http.StatusFound, founduser)

// 	}
// }

// func ProductViewerAdmin () gin.HandlerFunc{

// }

// func SearchProduct (id string) gin.HandlerFunc {

// }

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

// func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

// }
