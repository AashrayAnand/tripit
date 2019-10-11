package driver

import (
	"context"
	"log"
	"time"

	"github.com/AashrayAnand/Bill-List/models"
	"github.com/AashrayAnand/Bill-List/secret"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB client object
var Client = InitDB()

// users collection
var Users = Client.Database("BillList").Collection("users")

// initialize and return mongoDB client
func InitDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(secret.CONNSTRING))
	if err != nil {
		log.Fatal("init", err)
	}
	return client
}

func FindUser(user string, res *models.User) error {
	return Users.FindOne(context.TODO(), bson.D{{"username", user}}).Decode(res)
}

func AddUser(user string, pass string) error {
	var data models.User
	data.Id = string(uuid.New())
	data.Username = user
	data.Password = pass
	data.Created = time.Now()
	_, err := Users.InsertOne(context.TODO(), data)
	return err
}
