package main

import (
	"context"
	"log"

	"github.com/AashrayAnand/Bill-List/secret"
	"github.com/AashrayAnand/Bill-List/user"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB client object
var Client = InitDB()

// users collection
var users = Client.Database("BillList").Collection("users")

func main() {
	router := gin.Default() // initializ gin routing engine

	user.AddUserRoutes(router) // add all routes in the /user group
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.Run() // listen and server 0.0.0.0:8080
}

// initialize and return mongoDB client
func InitDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(secret.CONNSTRING))
	if err != nil {
		log.Fatal("init", err)
	}
	return client
}
