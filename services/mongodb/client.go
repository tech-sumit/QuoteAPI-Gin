package connections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"os"
	"time"
)

//MongoDB Initialisation
func MongoInit() (*mongo.Client, error) {
	var MongoHost = os.Getenv("MONGO_HOST")
	var MongoPort = os.Getenv("MONGO_PORT")
	var MongoUser = os.Getenv("MONGO_USER")
	var MongoCred = os.Getenv("MONGO_PASSWORD")
	var MongoUrl string
	if MongoUser != "" && MongoCred != "" {
		MongoUrl = "mongodb://" + MongoUser + ":" + MongoCred + "@" + MongoHost + ":" + MongoPort
	} else {
		MongoUrl = "mongodb://" + MongoHost + ":" + MongoPort
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoUrl))
	if err != nil {
		return nil, err
	}
	return client, nil
}
