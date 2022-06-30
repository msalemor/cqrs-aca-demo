package dataaccess

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func ConnectDB(mongoURI string) *mongo.Client {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//ping the database
// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Connected to MongoDB")
// 	return client
// }

// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func ConnectDB(mongoURI, database, collectionName string) *mongo.Collection {

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	collection := client.Database(database).Collection(collectionName)

	return collection
}
