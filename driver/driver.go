package driver

import (
	"context"
	"log"
	"time"

	"github.com/AashrayAnand/tripit/models"
	"github.com/AashrayAnand/tripit/secret"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB client object
var Client = InitDB()

// initialize pointer to mongodb users collection
var Users = Client.Database("BillList").Collection("users")

// TODO: POINTERS TO OTHER COLLECTIONS

// initialize and return mongoDB client
func InitDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(secret.CONNSTRING))
	if err != nil {
		log.Fatal("init", err)
	}
	return client
}

// find a user given a specified user name
func FindUser(user string, res *models.User) error {
	return Users.FindOne(context.TODO(), bson.D{{"username", user}}).Decode(res)
}

func GetId(user string) uuid.UUID {
	data := new(models.User)
	_ = Users.FindOne(context.TODO(), bson.D{{"username", user}}).Decode(&data)
	return data.Id
}

// add a new user to the users collection
func AddUser(user string, pass string) error {
	var data models.User
	data.Id = uuid.New()
	data.Username = user
	data.Password = pass
	data.Created = time.Now()
	_, err := Users.InsertOne(context.TODO(), data)
	return err
}
