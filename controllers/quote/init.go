package quote

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var QuotesCollection *mongo.Collection

//Collection Init
func Init(client *mongo.Client) {
	databaseName := os.Getenv("MONGO_DATABASE")
	collection := os.Getenv("MONGO_COLLECTION")
	database := client.Database(databaseName)
	QuotesCollection = database.Collection(collection)

	//Index creation
	_, _ = QuotesCollection.Indexes().CreateOne(context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"quote":  "text",
				"author": "text",
			},
			Options: &options.IndexOptions{
				Weights: map[string]int{
					"quote":  9,
					"author": 3,
				},
			},
		})
	log.Println("Database Initialised")
}
